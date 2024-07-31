package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresDB struct {
	Pool *pgxpool.Pool
}

func NewPostgresDB(connString string) (*PostgresDB, error) {
	pool, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}

	return &PostgresDB{Pool: pool}, nil
}

func (db *PostgresDB) Close() {
	db.Pool.Close()
}
