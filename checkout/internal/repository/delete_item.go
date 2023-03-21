package repository

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (r repository) DeleteItem(ctx context.Context, userId int64, sku uint32, count uint32) error {
	conn := r.connManager.GetQueryEngine(ctx)

	query, args, err := squirrel.
		Update("cart_item").
		Set("count", squirrel.Expr("GREATEST(count - ?, 0)", count)).
		Where("sku = ? and user_id =?", sku, userId).
		PlaceholderFormat(squirrel.Dollar).ToSql()

	if err != nil {
		return err
	}

	if _, err = conn.Exec(ctx, query, args...); err != nil {
		return errors.Wrap(err, "deleting item")
	}

	return nil
}
