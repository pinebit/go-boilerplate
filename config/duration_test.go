package config_test

import (
	"github.com/pinebit/go-boilerplate/config"
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	t.Parallel()
	d := config.MakeDuration(5 * time.Second)
	text, err := d.MarshalText()
	if err != nil || string(text) != "5s" {
		t.Fatalf("marshal: %q, %v", text, err)
	}
	var got config.Duration
	if err := got.UnmarshalText(text); err != nil {
		t.Fatal(err)
	}
	if got.Duration() != d.Duration() || got.String() != "5s" {
		t.Fatal("round trip mismatch")
	}
	if err := got.UnmarshalText([]byte("invalid")); err == nil {
		t.Fatal("expected invalid duration error")
	}
	if got.Duration() != 5*time.Second {
		t.Fatal("invalid input mutated duration")
	}
	if _, err := config.ParseDuration("invalid"); err == nil {
		t.Fatal("expected parse error")
	}
}
