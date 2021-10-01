package gocache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	cfg := NewConfig().WithAddr("127.0.0.1:6379").WithDB(0).WithExpiration(4)

	client, err := NewClient(cfg)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, client)
	assert.Equal(t, true, client.Ping())
	assert.Equal(t, nil, client.Close())
}

func TestRedisSetGet(t *testing.T) {
	cfg := NewConfig().WithAddr("127.0.0.1:6379").WithDB(0).WithExpiration(4)

	client, err := NewClient(cfg)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, client)

	data := "this is a test data"

	assert.Equal(t, nil, client.Set("test", data))
	assert.Equal(t, data, client.Get("test"))

	assert.Equal(t, nil, client.Close())
}

func TestMemcachedSetGet(t *testing.T) {
	cfg := NewConfig().WithDriver(MemcachedDriver).WithAddr("127.0.0.1:11211").WithDB(0).WithExpiration(4)

	client, err := NewClient(cfg)

	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, client)

	data := "this is a test data"

	assert.Equal(t, nil, client.Set("test", data))
	assert.Equal(t, data, client.Get("test"))

	assert.Equal(t, nil, client.Close())
}
