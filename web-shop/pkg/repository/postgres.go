package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

func NewPostgresDB(ctx context.Context, host string, port string, database string, user string, password string) (*pgx.Conn, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, database)

	return pgx.Connect(ctx, connString)
}
