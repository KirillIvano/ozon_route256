package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (r repository) ApplyOrderReservations(ctx context.Context, orderId int64) error {
	engine := r.connManager.GetQueryEngine(ctx)

	reservations, err := r.GetReservations(ctx, orderId)
	if err != nil {
		return errors.Wrap(err, "creating reservations")
	}

	for _, reservation := range reservations {
		query, args, err := sq.
			Update("warehouse_items").
			Set("count", sq.Expr("count-?", reservation.Count)).
			Where(sq.Eq{"sku": reservation.Sku}).
			Where(sq.Eq{"warehouse_id": reservation.WarehouseId}).
			PlaceholderFormat(sq.Dollar).
			ToSql()

		if err != nil {
			return err
		}

		_, err = engine.Exec(ctx, query, args...)

		if err != nil {
			return err
		}
	}

	return nil
}
