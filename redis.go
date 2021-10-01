package gocache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// redisCache client of the redis client
type redisCache struct {
	exp    int             // expiration of setted values
	client *redis.Client   // client
	ctx    context.Context // context
}

func newRedisCache(cfg *Config) (Client, error) {
	address := "localhost:6379"

	if len(cfg.Addresses) > 0 {
		address = cfg.Addresses[0]
	}

	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Username: cfg.Username,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return &redisCache{
		exp:    cfg.Expiration,
		client: client,
		ctx:    context.Background(),
	}, nil
}

// Gets from cache by key
func (rc *redisCache) Get(key string) interface{} {
	data, err := rc.client.Get(rc.ctx, key).Result()

	if err != nil {
		return nil
	}
	return data
}

// Sets to cache with key of data
// If something went, returns error
func (rc *redisCache) Set(key string, data interface{}) error {
	if err := rc.client.Set(rc.ctx, key, data, time.Duration(rc.exp)*time.Second).Err(); err != nil {
		return err
	}

	return nil
}
