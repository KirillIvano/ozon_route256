package repository

import (
	"context"
	"route256/loms/internal/domain"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/pkg/errors"
)

func (r repository) GetOutdatedOrders(ctx context.Context) ([]int64, error) {
	conn := r.connManager.GetQueryEngine(ctx)
	query, args, err := sq.
		Select("order_id").
		From("loms_order").
		Where("order_status = ? and now() > updated_at + interval '10 minutes'", domain.OrderStatusAwaitingPayment).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "forming get outdated orders query")
	}

	var orders []int64
	if err := pgxscan.Select(ctx, conn, &orders, query, args...); err != nil {
		return nil, errors.Wrap(err, "getting order info")
	}

	return orders, nil
}
