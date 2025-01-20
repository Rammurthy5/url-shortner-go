package config

import (
	"time"

	"github.com/Rammurthy5/url-shortner-go/internal/middleware"
	"github.com/jackc/pgx/v5"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/zap"
)

func InitDependencies() (Config, *zap.Logger, *pgx.Conn, *middleware.IdempotencyMiddleware, *sdktrace.TracerProvider) {
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
	// Initialize tracer
	tp, err := initTracer()
	if err != nil {
		logger.Fatal("Failed to initialize tracer", zap.Error(err))
	}
	tracerProvider = tp
	return config, logger, db, idempotencyMiddleware, tracerProvider
}
