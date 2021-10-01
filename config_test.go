package gocache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// newConfig returns default configurations
func newConfig() *Config {
	return NewConfig()
}

func TestNewConfig(t *testing.T) {
	cfg := newConfig()

	assert.Equal(t, RedisDriver, cfg.Driver)
	assert.Equal(t, []string{"127.0.0.1:6379"}, cfg.Addresses)
	assert.Equal(t, "", cfg.Username)
	assert.Equal(t, "", cfg.Password)
	assert.Equal(t, 0, cfg.DB)
	assert.Equal(t, 3, cfg.Expiration)
}

func TestWithDriver(t *testing.T) {
	cfg := newConfig().WithDriver(MemcachedDriver)

	assert.Equal(t, MemcachedDriver, cfg.Driver)
}

func TestWithAddr(t *testing.T) {
	cfg := newConfig().WithAddr("127.0.0.1:6379", "127.0.0.1:6377")

	assert.Equal(t, []string{"127.0.0.1:6379", "127.0.0.1:6377"}, cfg.Addresses)
}

func TestWithCreds(t *testing.T) {
	cfg := newConfig().WithCreds("user", "pass")

	assert.Equal(t, "user", cfg.Username)
	assert.Equal(t, "pass", cfg.Password)
}

func TestWithExpiration(t *testing.T) {
	cfg := newConfig().WithExpiration(10)

	assert.Equal(t, 10, cfg.Expiration)
}
