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
		_logger = zap.Must(zap.NewProduction())
	})
	return _logger
}

func ShutdownLogger() {
	if _logger != nil {
		err := _logger.Sync()
		if err != nil {
			log.Printf("Logger failed to sync: %v", err)
		}
	}
}
