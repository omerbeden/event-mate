package postgresadapter

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/omerbeden/event-mate/backend/tatooine/pkg/db"
)

type PgxAdapter struct {
	pool *pgxpool.Pool
}

func (p *PgxAdapter) Begin(ctx context.Context) (db.Tx, error) {
	tx, err := p.pool.Begin(ctx)
	if err != nil {
		return nil, err
	}

	return NewPgxTxAdapter(tx), nil

}

func (p *PgxAdapter) CopyFrom(ctx context.Context, tableName db.Identifier, columnNames []string, rowSrc db.CopyFromSource) (int64, error) {
	pgxIdentifier := pgx.Identifier(tableName)

	return p.pool.CopyFrom(ctx, pgxIdentifier, columnNames, rowSrc)
}

func (p *PgxAdapter) Exec(ctx context.Context, sql string, arguments ...any) (db.CommandTag, error) {
	commandTag, err := p.pool.Exec(ctx, sql, arguments...)

	if err != nil {
		return db.CommandTag{}, err
	}

	rowsAffected := commandTag.RowsAffected()
	return db.CommandTag{RowsAffected: rowsAffected}, nil
}

func (p *PgxAdapter) Query(ctx context.Context, sql string, args ...any) (db.Rows, error) {
	return p.pool.Query(ctx, sql, args...)
}

func (p *PgxAdapter) QueryRow(ctx context.Context, sql string, args ...any) db.Row {
	return p.pool.QueryRow(ctx, sql, args...)
}

func NewPgxAdapter(pool *pgxpool.Pool) *PgxAdapter {
	return &PgxAdapter{pool: pool}
}

type PgxTxAdapter struct {
	tx pgx.Tx
}

func (p *PgxTxAdapter) Commit(ctx context.Context) error {
	return p.tx.Commit(ctx)
}

func (p *PgxTxAdapter) CopyFrom(ctx context.Context, tableName db.Identifier, columnNames []string, rowSrc db.CopyFromSource) (int64, error) {
	pgxIdentifier := pgx.Identifier(tableName) // Direct conversion if both are slice of strings

	return p.tx.CopyFrom(ctx, pgxIdentifier, columnNames, rowSrc)
}

func (p *PgxTxAdapter) Exec(ctx context.Context, sql string, arguments ...any) (commandTag db.CommandTag, err error) {
	ct, err := p.tx.Exec(ctx, sql, arguments...)

	if err != nil {
		return db.CommandTag{}, err
	}

	rowsAffected := ct.RowsAffected()
	return db.CommandTag{RowsAffected: rowsAffected}, nil
}

func (p *PgxTxAdapter) Query(ctx context.Context, sql string, args ...any) (db.Rows, error) {
	return p.tx.Query(ctx, sql, args...)
}

func (p *PgxTxAdapter) QueryRow(ctx context.Context, sql string, args ...any) db.Row {
	return p.tx.QueryRow(ctx, sql, args...)
}

func (p *PgxTxAdapter) Rollback(ctx context.Context) error {
	return p.tx.Rollback(ctx)
}

func NewPgxTxAdapter(tx pgx.Tx) *PgxTxAdapter {
	return &PgxTxAdapter{tx: tx}
}
