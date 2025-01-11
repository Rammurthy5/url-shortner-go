package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"sync"
)

var (
	onceDB  sync.Once
	_dbInst *pgx.Conn
)

func GetDB(cfg Config) *pgx.Conn {

	onceDB.Do(func() {
		ctx := context.Background()
		conn_string := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
			cfg.DBConfig.Username, cfg.DBConfig.Password, cfg.DBConfig.Host, cfg.DBConfig.Port, cfg.DBConfig.Database)
		fmt.Println(conn_string)
		conn, err := pgx.Connect(ctx, conn_string)
		//conn, err := pgx.Connect(ctx, "")
		if err != nil {
			panic(fmt.Sprintf("Database failed to connect %s", err))
		}
		_dbInst = conn
	})
	return _dbInst
}

// CloseDB Function to gracefully close the DB connection
func ShutDownDB() {
	if _dbInst != nil {
		ctx := context.Background()
		_ = _dbInst.Close(ctx) // Handle the error gracefully
	}
}
