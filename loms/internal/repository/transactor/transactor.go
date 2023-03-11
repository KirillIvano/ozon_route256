package transactor

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type QueryEngine interface {
	Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
}

type TransactionManager struct {
	conn *pgx.Conn
}

type txkey string

const key = txkey("tx")

func (tm *TransactionManager) RunTransaction(ctx context.Context, level pgx.TxIsoLevel, fx func(ctxTX context.Context) error) error {
	tx, err := tm.conn.BeginTx(ctx,
		pgx.TxOptions{
			IsoLevel: pgx.RepeatableRead,
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

func New(conn *pgx.Conn) *TransactionManager {
	return &TransactionManager{
		conn: conn,
	}
}
