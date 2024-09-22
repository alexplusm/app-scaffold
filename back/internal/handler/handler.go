package handler

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

type handler struct {
	conn *pgx.Conn
}

func NewHandler(connDB *pgx.Conn) *handler {
	return &handler{
		conn: connDB,
	}
}

func (h *handler) Hello(c echo.Context) error {
	return c.String(200, "hello")
}

func (h *handler) Test(c echo.Context) error {
	var res string
	err := h.conn.QueryRow(context.Background(), "SELECT name FROM test LIMIT 1;").Scan(&res)

	if err != nil && !(err.Error() == pgx.ErrNoRows.Error()) {
		return err
	}
	return c.String(200, res)
}

func (h *handler) Close() error {
	return h.conn.Close(context.Background())
}
