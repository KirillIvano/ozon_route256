package repository

import (
	"context"
	"route256/libs/transactor"
	"route256/loms/internal/domain"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repository struct {
	connManager *transactor.TransactionManager
}

type LomsRepository interface {
	ListStocks(ctx context.Context, sku uint32) ([]domain.Stock, error)
	CreateOrder(ctx context.Context, userId domain.UserId) (int64, error)
	ClearReservations(ctx context.Context, orderId int64) error
	CreateOrderItems(ctx context.Context, orderId int64, items []domain.OrderItem) error
	GetOrderInfo(ctx context.Context, orderId int64) (*domain.OrderInfo, error)
	ApplyOrderReservations(ctx context.Context, orderId int64) error
	CreateReservation(ctx context.Context, reservations []domain.Reservation) error
	UpdateOrderStatus(ctx context.Context, orderId int64, status string) error
	RunReadCommitedTransaction(ctx context.Context, fx func(ctxTX context.Context) error) error
	GetOrderStatus(ctx context.Context, orderId int64) (string, error)
	GetOutdatedOrders(ctx context.Context) ([]int64, error)
}

var _ domain.LomsRepository = (*repository)(nil)

func Connect(conn *pgxpool.Pool) repository {
	return repository{
		connManager: transactor.New(conn),
	}
}
