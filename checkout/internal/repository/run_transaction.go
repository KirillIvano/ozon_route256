package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (r repository) RunReadCommitedTransaction(ctx context.Context, fx func(ctxTX context.Context) error) error {
	return r.connManager.RunTransaction(ctx, pgx.ReadCommitted, fx)
}
