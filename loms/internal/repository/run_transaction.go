package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// TODO: remove
func (r repository) RunTransaction(ctx context.Context, level pgx.TxIsoLevel, fx func(ctxTX context.Context) error) error {
	return r.ConnManager.RunTransaction(ctx, level, fx)
}
