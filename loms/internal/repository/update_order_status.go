package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (r repository) UpdateOrderStatus(ctx context.Context, orderId int64, status string) error {
	conn := r.ConnManager.GetQueryEngine(ctx)

	query, args, err := sq.Update("loms_order").
		Set("order_status", status).
		Where("order_id = ?", orderId).
		PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return errors.Wrap(err, "forming update status query")
	}

	_, err = conn.Exec(ctx, query, args...)

	return errors.Wrap(err, "updating status in db")
}
