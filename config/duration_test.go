package config_test

import (
	"testing"
	"time"

	"github.com/pinebit/go-boilerplate/config"
	"github.com/stretchr/testify/assert"
)

func TestDuration(t *testing.T) {
	t.Parallel()

	d := config.MakeDuration(5 * time.Second)
	assert.Equal(t, 5*time.Second, d.Duration())
	assert.Equal(t, "5s", d.String())

	d, err := config.ParseDuration("3s")
	assert.NoError(t, err)
	assert.Equal(t, 3*time.Second, d.Duration())

	_, err = config.ParseDuration("xxx")
	assert.Error(t, err)
}
