package config

import (
	"go.uber.org/zap"
	"log"
	"sync"
)

var (
	once    sync.Once
	_logger *zap.Logger
)

func GetLogger() *zap.Logger {
	once.Do(func() {
		baseLogger := zap.Must(zap.NewProduction())
		defer func(logger *zap.Logger) {
			err := logger.Sync()
			if err != nil {
				log.Fatal("Logger failed to sync", zap.Error(err))
			}
		}(baseLogger)
	})
	return _logger
}
