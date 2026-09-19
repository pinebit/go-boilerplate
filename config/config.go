package config

import (
	"bytes"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"time"

	"github.com/pelletier/go-toml/v2"
)

type HttpServerConfig struct {
	Address           string
	Port              uint16
	ReadHeaderTimeout Duration
	ReadTimeout       Duration
	WriteTimeout      Duration
	IdleTimeout       Duration
}

type Config struct {
	DevMode         bool
	ShutdownTimeout Duration
	HttpServer      *HttpServerConfig
}

func NewDefaultConfig() *Config {
	return &Config{DevMode: true, ShutdownTimeout: MakeDuration(5 * time.Second), HttpServer: &HttpServerConfig{
		Port: 8080, ReadHeaderTimeout: MakeDuration(5 * time.Second), ReadTimeout: MakeDuration(15 * time.Second),
		WriteTimeout: MakeDuration(30 * time.Second), IdleTimeout: MakeDuration(60 * time.Second),
	}}
}

func (c *Config) Validate() error {
	if c.HttpServer == nil {
		return errors.New("HttpServer configuration is required")
	}
	if c.HttpServer.Port == 0 {
		return errors.New("HttpServer.Port must be greater than zero")
	}
	for _, value := range []struct {
		name     string
		duration Duration
	}{
		{"ShutdownTimeout", c.ShutdownTimeout}, {"HttpServer.ReadHeaderTimeout", c.HttpServer.ReadHeaderTimeout},
		{"HttpServer.ReadTimeout", c.HttpServer.ReadTimeout}, {"HttpServer.WriteTimeout", c.HttpServer.WriteTimeout},
		{"HttpServer.IdleTimeout", c.HttpServer.IdleTimeout},
	} {
		if value.duration.Duration() <= 0 {
			return fmt.Errorf("%s must be positive", value.name)
		}
	}
	return nil
}

func (c *Config) LoadFromToml(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read config %q: %w", path, err)
	}
	// Decode a copy so invalid configuration cannot partially mutate the receiver.
	next := *c
	if c.HttpServer != nil {
		server := *c.HttpServer
		next.HttpServer = &server
	}
	if err := toml.NewDecoder(bytes.NewReader(data)).DisallowUnknownFields().Decode(&next); err != nil {
		return fmt.Errorf("decode config %q: %w", path, err)
	}
	if err := next.Validate(); err != nil {
		return fmt.Errorf("validate config %q: %w", path, err)
	}
	*c = next
	return nil
}

func (c Config) WriteToml(path string, perm fs.FileMode) error {
	if err := c.Validate(); err != nil {
		return err
	}
	var buf bytes.Buffer
	enc := toml.NewEncoder(&buf)
	enc.SetIndentTables(true)
	if err := enc.Encode(c); err != nil {
		return err
	}
	return os.WriteFile(path, buf.Bytes(), perm)
}
