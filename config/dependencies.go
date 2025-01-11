package config

import (
	"github.com/Rammurthy5/url-shortner-go/internal/middleware"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
	"time"
)

func InitDependencies() (Config, *zap.Logger, *pgx.Conn, *middleware.IdempotencyMiddleware) {
	config, err := Load()
	if err != nil {
		panic(err)
	}
	logger := GetLogger()
	db := GetDB(config)
	cache := GetCache(config)
	idempotencyMiddleware := &middleware.IdempotencyMiddleware{
		RedisClient: cache,
		TTL:         10 * time.Minute, // Set TTL for idempotency keys
	}
	return config, logger, db, idempotencyMiddleware
}
