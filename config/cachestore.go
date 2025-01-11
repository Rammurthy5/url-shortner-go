package config

import (
	"github.com/go-redis/redis/v8"
	"sync"
)

var (
	onceCacheStore sync.Once
	_cache         *redis.Client
)

func GetCache(cfg Config) *redis.Client {
	onceCacheStore.Do(func() {
		// Initialize Redis client
		redisClient := redis.NewClient(&redis.Options{
			Addr:     cfg.CacheConfig.Host,
			Password: cfg.CacheConfig.Password,
			DB:       cfg.CacheConfig.Db,
		})
		_cache = redisClient
	})
	return _cache
}

func ShutDownCache() {
	if _cache != nil {
		_ = _cache.Close()
	}
}
