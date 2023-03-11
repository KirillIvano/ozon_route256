package repository

import (
	"context"
	"route256/loms/internal/domain"
	"route256/loms/internal/repository/schema"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/pkg/errors"
)

func (r repository) GetOrderInfo(ctx context.Context, orderId int64) (*domain.OrderInfo, error) {
	conn := r.connManager.GetQueryEngine(ctx)
	query, args, err := sq.
		Select("user_id", "order_id", "order_status", "created_at").
		From("loms_order").
		Where(sq.Eq{"order_id": orderId}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "forming order info request")
	}

	var order schema.Order
	if err := pgxscan.Get(ctx, conn, &order, query, args...); err != nil {
		return nil, errors.Wrap(err, "getting order info")
	}

	orderItems, err := r.GetOrderItems(ctx, orderId)
	if err != nil {
		return nil, errors.Wrap(err, "getting order items")
	}

	return &domain.OrderInfo{
		Status: order.OrderStatus,
		User:   order.UserId,
		Items:  orderItems,
	}, nil
}
