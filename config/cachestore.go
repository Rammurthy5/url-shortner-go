package config

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
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
		_, err := redisClient.Ping(context.Background()).Result()
		if err != nil {
			log.Fatalf("Failed to connect to Redis: %v", err)
		}
		_cache = redisClient
	})
	return _cache
}

func ShutDownCache() {
	if _cache != nil {
		_ = _cache.Close()
	}
}
