package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresConfig struct {
	pgxpool.Config
	ConnectionString string
}

func NewConn(config *PostgresConfig) *pgxpool.Pool {
	configFrom, err := pgxpool.ParseConfig(config.ConnectionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to parse config: %v\n", err)
		os.Exit(1)
	}
	configFrom.MinConns = config.MinConns
	configFrom.MaxConns = config.MaxConns

	//later import db tracer

	pool, err := pgxpool.NewWithConfig(context.Background(), configFrom)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	return pool
}
