package db

import (
	"context"
)

type Executor interface {
	Begin(ctx context.Context) (Tx, error)
	CopyFrom(ctx context.Context, tableName Identifier, columnNames []string, rowSrc CopyFromSource) (int64, error)
	Exec(ctx context.Context, sql string, arguments ...any) (CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) Row
}

type Tx interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	CopyFrom(ctx context.Context, tableName Identifier, columnNames []string, rowSrc CopyFromSource) (int64, error)
	Exec(ctx context.Context, sql string, arguments ...any) (commandTag CommandTag, err error)
	Query(ctx context.Context, sql string, args ...any) (Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) Row
}

type Identifier []string

type CommandTag struct {
	RowsAffected int64
}
type CopyFromSource interface {
	Next() bool
	Values() ([]any, error)
	Err() error
}
type Row interface {
	Scan(dest ...any) error
}
type Rows interface {
	Close()
	Err() error
	Next() bool
	Scan(dest ...any) error
}

type TransactionManager interface {
	Begin(ctx context.Context) (Tx, error)
}
