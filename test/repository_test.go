package test

import (
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

func initRepo() (*pgxpool.Pool, error) {

}

func TestMain(m *testing.M) {
	var conn *pgxpool.Pool
	var err error
	repo, conn, err = initRepo()
	defer conn.Close()
	if err != nil {
		log.Panicf("Failed to initalize repositry: %v", err)
	}
	code := m.Run()

	conn.Exec(ctx,
		` DROP TABLE case_files;
			DROP TABLE agencies;
	`)
	os.Exit(code)
}
