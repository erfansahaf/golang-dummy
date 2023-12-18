package database

import (
	"database/sql"
	"fmt"
	"time"
)

type DBConfig struct {
	Username       string
	Password       string
	DBName         string
	MaxConLifetime time.Duration
	MaxOpenIdleCon int
}

func InitDB(config DBConfig) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", config.Username, config.Password, config.DBName))
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(config.MaxConLifetime)
	db.SetMaxOpenConns(config.MaxOpenIdleCon)
	db.SetMaxIdleConns(config.MaxOpenIdleCon)

	return db
}
