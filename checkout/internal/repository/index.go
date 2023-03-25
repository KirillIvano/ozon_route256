package repository

import (
	"context"
	"route256/checkout/internal/domain"
	"route256/libs/transactor"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	connManager *transactor.TransactionManager
}

type CheckoutRepository interface {
	GetCartItems(ctx context.Context, userId int64) ([]domain.CartItem, error)
	DeleteItem(ctx context.Context, userId int64, sku uint32, count uint32) error
	DeleteCart(ctx context.Context, userId int64) error
	GetCartItemExists(ctx context.Context, userId int64, sku uint32) (bool, error)
	AddToCart(ctx context.Context, userId int64, sku uint32, count uint32) error
}

var _ domain.CheckoutRepository = (*repository)(nil)

func New(conn *pgxpool.Pool) repository {
	return repository{
		connManager: transactor.New(conn),
	}
}
