package repository

import (
	"context"
	"route256/loms/internal/domain"
	"route256/loms/internal/repository/schema"

	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/pkg/errors"
)

func (r repository) ListStocks(ctx context.Context, sku uint32) ([]domain.Stock, error) {
	conn := r.connManager.GetQueryEngine(ctx)

	query, args, err := sq.Select("warehouse_items.warehouse_id", "warehouse_items.count - SUM(COALESCE(reservation.count, 0)) AS count").
		From("warehouse_items").
		LeftJoin("reservation").
		JoinClause("on warehouse_items.warehouse_id = reservation.warehouse_id and warehouse_items.sku = reservation.sku").
		Where("warehouse_items.sku = ?", sku).
		GroupBy("warehouse_items.warehouse_id", "warehouse_items.sku").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "formatting list stocks query failed")
	}

	var items []schema.WarehouseItem
	err = pgxscan.Select(ctx, conn, &items, query, args...)

	if err != nil {
		return nil, errors.Wrap(err, "querying stocks failed")
	}

	stocks := make([]domain.Stock, len(items))
	for idx, item := range items {
		stocks[idx] = domain.Stock{
			WarehouseID: item.WarehouseId,
			Count:       uint64(item.Count),
		}
	}

	return stocks, err
}
