package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type ConnectionDB struct {
	conn *pgx.Conn
}

func NewConnection(url string) (connDB *ConnectionDB, err error) {
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}

	return &ConnectionDB{
		conn: conn,
	}, nil
}

func (c *ConnectionDB) GetTestName() ([]string, error) {
	var res []string
	rows, err := c.conn.Query(context.Background(), "SELECT name FROM test;")
	if err != nil && !(err.Error() == pgx.ErrNoRows.Error()) {
		return res, err
	}

	var name string
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			return res, err
		}
		res = append(res, name)
	}

	return res, nil
}

func (c *ConnectionDB) NewTestName(name string) error {
	createTime := time.Now()
	_, err := c.conn.Exec(context.Background(), "INSERT INTO test (name, create_time) VALUES ($1, $2)", name, createTime)
	if err != nil {
		return err
	}

	return nil

}

func (c *ConnectionDB) Close() error {
	return c.conn.Close(context.Background())
}
