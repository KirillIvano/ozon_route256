package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

func (r repository) createCartItem(ctx context.Context, userId int64, sku uint32, count uint32) error {
	conn := r.connManager.GetQueryEngine(ctx)

	query, args, err := sq.
		Insert("cart_item").
		Columns("user_id", "sku", "count").
		Values(userId, sku, count).PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, query, args...)

	return err
}

func (r repository) updateItemsCount(ctx context.Context, userId int64, sku uint32, count uint32) error {
	conn := r.connManager.GetQueryEngine(ctx)

	query, args, err := sq.
		Update("cart_item").
		Set("count", sq.Expr("count + ?", count)).
		Where("sku = ? and user_id =?", sku, userId).
		PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, query, args...)

	return err
}

func (r repository) AddToCart(ctx context.Context, userId int64, sku uint32, count uint32) error {
	err := r.RunReadCommitedTransaction(ctx, func(ctxTX context.Context) error {
		exists, err := r.GetCartItemExists(ctx, userId, sku)
		if err != nil {
			return errors.Wrap(err, "checking for existing")
		}

		if exists {
			err := r.updateItemsCount(ctx, userId, sku, count)
			if err != nil {
				return errors.Wrap(err, "updating cart item count")
			}
		} else {
			err := r.createCartItem(ctx, userId, sku, count)
			if err != nil {
				return errors.Wrap(err, "creating cart itme")
			}
		}

		return nil
	})

	return err
}
