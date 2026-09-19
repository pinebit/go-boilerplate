package config_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/pinebit/go-boilerplate/config"
)

func TestConfigRoundTrip(t *testing.T) {
	t.Parallel()
	cfg := config.NewDefaultConfig()
	if err := cfg.Validate(); err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(t.TempDir(), "config.toml")
	if err := cfg.WriteToml(path, 0600); err != nil {
		t.Fatal(err)
	}
	got := config.NewDefaultConfig()
	if err := got.LoadFromToml(path); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, cfg) {
		t.Fatalf("round trip: got %+v, want %+v", got, cfg)
	}
}

func TestLoadConfig(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		name, input string
		invalid     bool
	}{
		{"partial", "[HttpServer]\nPort = 3333", false},
		{"unknown", "DevMod = true", true},
		{"invalid duration", "ShutdownTimeout = 'bad'", true},
		{"negative duration", "ShutdownTimeout = '-1s'", true},
		{"zero port", "[HttpServer]\nPort = 0", true},
		{"invalid syntax", "[", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			path := filepath.Join(t.TempDir(), "config.toml")
			if err := os.WriteFile(path, []byte(tc.input), 0600); err != nil {
				t.Fatal(err)
			}
			cfg := config.NewDefaultConfig()
			err := cfg.LoadFromToml(path)
			if (err != nil) != tc.invalid {
				t.Fatalf("unexpected error: %v", err)
			}
			if tc.invalid && !reflect.DeepEqual(cfg, config.NewDefaultConfig()) {
				t.Fatal("invalid input mutated configuration")
			}
			if !tc.invalid && (cfg.HttpServer.Port != 3333 || cfg.HttpServer.ReadTimeout.Duration() <= 0) {
				t.Fatal("partial configuration lost defaults")
			}
		})
	}
}

func TestMissingConfig(t *testing.T) {
	if err := config.NewDefaultConfig().LoadFromToml(filepath.Join(t.TempDir(), "missing")); err == nil {
		t.Fatal("expected error")
	}
}
