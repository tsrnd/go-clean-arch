package cache

import "time"

// Cache interface
type Cache interface {
	Get(key string) (string, error)
	Set(key, value string, expiration time.Duration) error
}
