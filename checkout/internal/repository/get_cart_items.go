package repository

import (
	"context"
	"route256/checkout/internal/domain"
	"route256/checkout/internal/repository/schema"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
)

func (r repository) GetCartItems(ctx context.Context, userId int64) ([]domain.CartItem, error) {
	conn := r.connManager.GetQueryEngine(ctx)

	query, args, err := sq.
		Select("user_id", "sku", "count").
		From("cart_item").
		Where("user_id =?", userId).
		Where("count > 0").
		PlaceholderFormat(sq.Dollar).ToSql()

	if err != nil {
		return nil, err
	}

	var scannedRes []schema.CartItem
	if err = pgxscan.Select(ctx, conn, &scannedRes, query, args...); err != nil {
		return nil, err
	}

	domainRes := make([]domain.CartItem, len(scannedRes))
	for i, scannedItem := range scannedRes {
		domainRes[i] = domain.CartItem{
			UserId: uint32(userId),
			Count:  uint32(scannedItem.Count),
			Sku:    uint32(scannedItem.Sku),
		}
	}

	return domainRes, nil
}
