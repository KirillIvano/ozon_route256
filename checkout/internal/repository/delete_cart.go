package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (r repository) DeleteCart(ctx context.Context, userId int64) error {
	conn := r.connManager.GetQueryEngine(ctx)

	query, args, err := sq.Delete("cart_item").Where("user_id=?", userId).PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return errors.Wrap(err, "forming cart delete query")
	}

	if _, err := conn.Exec(ctx, query, args...); err != nil {
		return errors.Wrap(err, "deleting cart")
	}

	return nil
}
