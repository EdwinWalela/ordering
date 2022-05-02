package repository

import (
	"context"

	c "edwinwalela/ordering/config"

	"github.com/jackc/pgx/v4"
)

type Repository struct {
	Conn       *pgx.Conn
	Ctx        context.Context
	SmsService c.ATSMS
}
