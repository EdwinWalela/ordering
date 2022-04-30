package repository

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type Repository struct {
	Conn *pgx.Conn
	Ctx  context.Context
}
