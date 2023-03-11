package repository

import (
	"context"
	"route256/loms/internal/domain"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (r repository) CreateReservation(ctx context.Context, reservations []domain.Reservation) error {
	engine := r.ConnManager.GetQueryEngine(ctx)

	sql := sq.
		Insert("reservation").
		Columns("warehouse_id", "order_id", "sku", "count")

	for _, reserv := range reservations {
		sql = sql.Values(reserv.WarehouseId, reserv.OrderId, reserv.Sku, reserv.Count)
	}

	query, args, err := sql.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = engine.Exec(ctx, query, args...)
	if err != nil {
		return errors.Wrap(err, "failed creating reservations")
	}

	return nil
}
