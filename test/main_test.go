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

func TestMain(m *testing.M) {
	ctx := context.Background()
	if err := godotenv.Load("../.env"); err != nil {
		if os.Getenv("PG_URL_LOCAL") == "" {
			log.Fatalf("Failed to initalize repository: %v", err)
		}
	}
	conn, err := pgx.Connect(ctx, os.Getenv("DB_URL"))
	defer conn.Close(ctx)

	if err != nil {
		log.Fatalf("Failed to initalize repository: %v", err)
	}
	r = repo.Repository{
		Conn: conn,
		Ctx:  ctx,
	}

	if err != nil {
		log.Fatalf("Failed to initalize repositry: %v", err)
	}
	code := m.Run()

	conn.Exec(ctx,
		` DROP TABLE case_files;
			DROP TABLE agencies;
	`)
	os.Exit(code)
}
