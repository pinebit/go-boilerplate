package config

import (
	"bytes"
	"io/fs"
	"os"
	"time"

	"github.com/pelletier/go-toml/v2"
)

type HttpServerConfig struct {
	Address string
	Port    uint16
}

type Config struct {
	DevMode         bool
	ShutdownTimeout Duration
	HttpServer      *HttpServerConfig
}

func NewDefaultConfig() *Config {
	return &Config{
		DevMode:         true,
		ShutdownTimeout: MakeDuration(5 * time.Second),
		HttpServer: &HttpServerConfig{
			Port: 8080,
		},
	}
}

func (c *Config) LoadFromToml(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	return toml.Unmarshal(data, c)
}

func (c Config) WriteToml(path string, perm fs.FileMode) error {
	buf := bytes.Buffer{}
	enc := toml.NewEncoder(&buf)
	enc.SetIndentTables(true)
	if err := enc.Encode(c); err != nil {
		return err
	}
	return os.WriteFile(path, buf.Bytes(), perm)
}
