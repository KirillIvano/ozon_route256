package repository

import (
	"context"
	"route256/loms/internal/domain"
	"route256/loms/internal/repository/schema"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (r repository) GetOrderItems(ctx context.Context, orderId int64) ([]domain.OrderItem, error) {
	conn := r.connManager.GetQueryEngine(ctx)
	query, args, err := sq.
		Select("sku", "count").
		From("order_items").
		Where(sq.Eq{"order_id": orderId}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, err
	}

	var orderItems []schema.OrderItem
	if err := pgxscan.Select(ctx, conn, &orderItems, query, args...); err != nil {
		return nil, err
	}

	domainItems := make([]domain.OrderItem, len(orderItems))
	for i, orderItem := range orderItems {
		domainItems[i] = domain.OrderItem{
			Sku:   uint32(orderItem.Sku),
			Count: uint32(orderItem.Count),
		}
	}

	return domainItems, nil
}
