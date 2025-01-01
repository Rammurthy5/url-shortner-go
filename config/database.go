package config

import (
	"database/sql"
	"fmt"
	"sync"
)

var (
	onceDb  sync.Once
	_dbInst *sql.DB
)

func GetDB(cfg Config) *sql.DB {
	onceDb.Do(func() {
		_, err := sql.Open("postgres",
			fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=mydatabase",
				cfg.DBConfig.Host, cfg.DBConfig.Port, cfg.DBConfig.Username, cfg.DBConfig.Password))
		if err != nil {
			panic("failed to connect to database")
		}
	})
	return _dbInst
}
