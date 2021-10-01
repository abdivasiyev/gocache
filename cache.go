package gocache

import "errors"

type Client interface {
	Ping() bool
	Get(key string) string
	Set(key string, data string) error
	Close() error
}

func NewClient(cfg *Config) (Client, error) {
	driver := cfg.Driver

	if driver == RedisDriver {
		return newRedisCache(cfg)
	}

	if driver == MemcachedDriver {
		return newMemCached(cfg)
	}

	return nil, errors.New("unsupported driver type")
}
