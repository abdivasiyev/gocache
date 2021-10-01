package gocache

import "errors"

type Client interface {
	Get(key string) interface{}
	Set(key string, data interface{}) error
}

func NewClient(cfg *Config) (Client, error) {
	driver := cfg.Driver

	if driver == RedisDriver {
		return newRedisCache(cfg)
	}

	return nil, errors.New("unsupported driver type")
}
