package repository

import (
	"context"
	"route256/loms/internal/domain"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/pkg/errors"
)

// add transaction for inserting in second table with items
func (r repository) CreateOrder(ctx context.Context, userId int64) (int64, error) {
	conn := r.connManager.GetQueryEngine(ctx)

	query, args, err := sq.Insert("loms_order").
		Columns("user_id", "order_status", "created_at").
		Values(userId, domain.OrderStatusNew, time.Now()).
		Suffix("RETURNING order_id").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return 0, errors.Wrap(err, "failed to create database query")
	}

	var id int64
	err = conn.QueryRow(
		ctx,
		query,
		args...,
	).Scan(&id)

	if err != nil {
		return 0, errors.Wrap(err, "creating order")
	}

	return id, nil
}
