package repository

import (
	"context"
	"route256/loms/internal/repository/schema"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (r repository) GetReservations(ctx context.Context, orderId int64) ([]schema.Reservation, error) {
	conn := r.connManager.GetQueryEngine(ctx)
	query, args, err := sq.
		Select("warehouse_id", "order_id", "sku", "count").
		From("reservation").
		Where(sq.Eq{"order_id": orderId}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	var reservations []schema.Reservation
	if err := pgxscan.Select(ctx, conn, &reservations, query, args...); err != nil {
		return nil, err
	}

	return reservations, nil
}
