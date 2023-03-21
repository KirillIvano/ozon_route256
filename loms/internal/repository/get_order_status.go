package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (r repository) GetOrderStatus(ctx context.Context, orderId int64) (string, error) {
	conn := r.connManager.GetQueryEngine(ctx)
	query, args, err := sq.
		Select("order_status").
		From("loms_order").
		Where(sq.Eq{"order_id": orderId}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return "", err
	}

	var status string
	if err := pgxscan.Get(ctx, conn, &status, query, args...); err != nil {
		return "", err
	}

	return status, nil
}
