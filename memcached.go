package gocache

import (
	"github.com/bradfitz/gomemcache/memcache"
)

// memCache client of the memcached client
type memCached struct {
	exp    int              // expiration of setted values
	client *memcache.Client // client
}

func newMemCached(cfg *Config) (Client, error) {

	client := memcache.New(cfg.Addresses...)

	return &memCached{
		exp:    cfg.Expiration,
		client: client,
	}, nil
}

// Ping checks cache alive
func (mc *memCached) Ping() bool {
	return mc.client.Ping() == nil
}

// Gets from cache by key
func (mc *memCached) Get(key string) string {
	data, err := mc.client.Get(key)

	if err != nil {
		return ""
	}

	return string(data.Value)
}

// Sets to cache with key of data
// If something went, returns error
func (rc *memCached) Set(key string, data string) error {
	item := &memcache.Item{
		Key:        key,
		Value:      []byte(data),
		Expiration: int32(rc.exp),
	}

	return rc.client.Set(item)
}

// Close closes connection
func (rc *memCached) Close() error {
	return nil
}
