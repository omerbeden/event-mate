package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresExecutor interface {
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

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
