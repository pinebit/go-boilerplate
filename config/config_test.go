package config_test

import (
	"os"
	"testing"

	"github.com/pinebit/go-boilerplate/config"

	"github.com/stretchr/testify/assert"
)

func TestNewDefaultConfig(t *testing.T) {
	t.Parallel()

	c := config.NewDefaultConfig()
	assert.True(t, c.DevMode)
	assert.NotNil(t, c.HttpServer)
	assert.Equal(t, 8080, int(c.HttpServer.Port))
	assert.Empty(t, c.HttpServer.Address)
}

func TestLoadFromToml(t *testing.T) {
	t.Parallel()

	c := config.NewDefaultConfig()
	err := c.LoadFromToml("testdata/sample.toml")
	assert.NoError(t, err)
	assert.NotNil(t, c.HttpServer)
	assert.Equal(t, 3333, int(c.HttpServer.Port))
}

func TestWriteToml(t *testing.T) {
	t.Parallel()

	const outputFile = "testdata/output.toml"
	c := config.NewDefaultConfig()
	err := c.WriteToml(outputFile, 0666)
	assert.NoError(t, err)
	err = c.LoadFromToml(outputFile)
	assert.NoError(t, err)
	assert.Equal(t, config.NewDefaultConfig(), c)
	os.Remove(outputFile)
}
