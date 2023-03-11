package repository

import (
	"context"
	"route256/loms/internal/domain"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (r repository) CreateOrderItems(ctx context.Context, orderId int64, items []domain.OrderItem) error {
	engine := r.ConnManager.GetQueryEngine(ctx)

	sql := sq.
		Insert("order_items").
		Columns("order_id", "sku", "count")

	for _, item := range items {
		sql = sql.Values(orderId, item.Sku, item.Count)
	}

	query, args, err := sql.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}

	_, err = engine.Exec(ctx, query, args...)
	if err != nil {
		return errors.Wrap(err, "failed creating order items")
	}

	return nil
}
