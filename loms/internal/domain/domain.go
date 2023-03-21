package domain

import (
	"context"
)

type Validator interface {
	Validate() error
}

type LomsRepository interface {
	ListStocks(ctx context.Context, sku uint32) ([]Stock, error)
	CreateOrder(ctx context.Context, userId UserId) (int64, error)
	ClearReservations(ctx context.Context, orderId int64) error
	CreateOrderItems(ctx context.Context, orderId int64, items []OrderItem) error
	GetOrderInfo(ctx context.Context, orderId int64) (*OrderInfo, error)
	ApplyOrderReservations(ctx context.Context, orderId int64) error
	CreateReservation(ctx context.Context, reservations []Reservation) error
	UpdateOrderStatus(ctx context.Context, orderId int64, status string) error
	RunReadCommitedTransaction(ctx context.Context, fx func(ctxTX context.Context) error) error
	GetOrderStatus(ctx context.Context, orderId int64) (string, error)
	GetOutdatedOrders(ctx context.Context) ([]int64, error)
}

type LomsDomain struct {
	lomsRepository LomsRepository
}

func New(lomsRepository LomsRepository) *LomsDomain {
	return &LomsDomain{
		lomsRepository: lomsRepository,
	}
}
