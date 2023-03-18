package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (r repository) UpdateOrderStatus(ctx context.Context, orderId int64, status string) error {
	conn := r.connManager.GetQueryEngine(ctx)

	query, args, err := sq.Update("loms_order").
		SetMap(map[string]any{
			"order_status": status,
			"updated_at":   sq.Expr("now()"),
		}).
		Where("order_id = ?", orderId).
		PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return errors.Wrap(err, "forming update status query")
	}

	_, err = conn.Exec(ctx, query, args...)

	return errors.Wrap(err, "updating status in db")
}
