package config

import (
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

func InitDependencies() (Config, *zap.Logger, *pgx.Conn) {
	config, err := Load()
	if err != nil {
		panic(err)
	}
	logger := GetLogger()
	db := GetDB(config)
	return config, logger, db
}
