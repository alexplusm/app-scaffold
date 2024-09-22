package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func NewConnection(url string) (conn *pgx.Conn, err error) {
	conn, err = pgx.Connect(context.Background(), url)

	return conn, err
}
