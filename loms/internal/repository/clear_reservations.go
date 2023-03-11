package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (r repository) ClearReservations(ctx context.Context, orderId int64) error {
	conn := r.ConnManager.GetQueryEngine(ctx)

	query, args, err := sq.Delete("reservation").
		Where("order_id = ?", orderId).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return errors.Wrap(err, "forming clear reservations query")
	}

	_, err = conn.Exec(ctx, query, args...)

	return errors.Wrap(err, "clearing reservations")
}
