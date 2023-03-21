package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (r repository) GetCartItemExists(ctx context.Context, userId int64, sku uint32) (bool, error) {
	conn := r.connManager.GetQueryEngine(ctx)

	query, args, err := sq.
		Select("COUNT(*) as cnt").
		From("cart_item").
		Where("user_id =? and sku = ?", userId, sku).
		PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return false, err
	}
	var cnt uint32
	if err = pgxscan.Get(ctx, conn, &cnt, query, args...); err != nil {
		return false, err
	}

	return cnt > 0, nil
}
