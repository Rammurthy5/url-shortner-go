package config

import (
	"database/sql"
	"go.uber.org/zap"
)

func InitDependencies() (Config, *zap.Logger, *sql.DB) {
	config, err := Load()
	if err != nil {
		panic(err)
	}
	logger := GetLogger()
	db := GetDB(config)
	return config, logger, db
}
