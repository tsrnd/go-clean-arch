package config

import (
	"os"

	"github.com/tsrnd/go-clean-arch/services/cache"
	"github.com/tsrnd/go-clean-arch/services/cache/redis"
)

// Cache func
func Cache() cache.Cache {
	return redis.Connect(
		os.Getenv("REDIS_ADDR"),
		os.Getenv("REDIS_PASSWORD"),
		0,
	)
}
