package test

import (
	"context"
	repo "edwinwalela/ordering/repository"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

var r repo.Repository

func initDb(conn *pgx.Conn) error {
	rawSqlBytes, err := os.ReadFile("../db.sql")
	if err != nil {
		return err
	}
	sql := string(rawSqlBytes)
	_, err = conn.Exec(context.Background(), sql)
	return err
}

func destroyDb(conn *pgx.Conn) error {
	_, err := conn.Exec(context.Background(), `
	DROP TABLE orders;
	DROP TABLE customers;
	`)
	return err
}

func TestMain(m *testing.M) {
	ctx := context.Background()
	if err := godotenv.Load("../.env"); err != nil {
		if os.Getenv("DB_URL") == "" {
			log.Fatalf("Failed to initalize repository: %v", err)
		}
	}
	conn, err := pgx.Connect(ctx, os.Getenv("DB_URL"))
	defer conn.Close(ctx)
	if err != nil {
		log.Fatalf("Failed to initalize repository: %v", err)
	}

	initDb(conn)
	r = repo.Repository{
		Conn: conn,
		Ctx:  ctx,
	}

	if err != nil {
		log.Fatalf("Failed to initalize repositry: %v", err)
	}
	code := m.Run()
	destroyDb(conn)
	os.Exit(code)
}
