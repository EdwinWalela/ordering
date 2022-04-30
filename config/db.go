package config

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

type DB struct {
	Conn *pgx.Conn
}

func (*DB) InitDb(cfg *Config) *DB {
	conn, err := pgx.Connect(context.Background(), cfg.DbURl)
	if err != nil {
		log.Panicf("Failed to connect to DB: %v", err)
	}
	return &DB{
		Conn: conn,
	}
}
