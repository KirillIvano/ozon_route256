package transactor

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type QueryEngine interface {
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
}

type TransactionManager struct {
	conn *pgxpool.Pool
}

type TxKey string

const key = TxKey("__txManager__")

func (tm *TransactionManager) RunTransaction(ctx context.Context, level pgx.TxIsoLevel, fx func(ctxTX context.Context) error) error {
	tx, err := tm.conn.BeginTx(ctx,
		pgx.TxOptions{
			IsoLevel: level,
		})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if err := fx(context.WithValue(ctx, key, tx)); err != nil {
		return err
	}

	err = tx.Commit(ctx)

	return err
}

func (tm *TransactionManager) GetQueryEngine(ctx context.Context) QueryEngine {
	tx, ok := ctx.Value(key).(QueryEngine)
	if ok && tx != nil {
		return tx
	}

	return tm.conn
}

func New(conn *pgxpool.Pool) *TransactionManager {
	return &TransactionManager{
		conn: conn,
	}
}
