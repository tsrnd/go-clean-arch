package redis

import (
	"time"

	"github.com/go-redis/redis"
	"github.com/tsrnd/go-clean-arch/services/cache"
)

// Redis struct
type Redis struct {
	c *redis.Client
}

// Connect func
func Connect(addr, password string, db int) cache.Cache {
	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &Redis{c: c}
}

// Get func
func (r *Redis) Get(key string) (string, error) {
	return r.c.Get(key).Result()
}

// Set func
func (r *Redis) Set(key, value string, expiration time.Duration) error {
	return r.c.Set(key, value, expiration).Err()
}
