package loms_server

import (
	"context"
	"route256/loms/internal/domain"
	lomsService "route256/loms/pkg/loms_service"

	"github.com/pkg/errors"
)

var (
	ErrEmptyUser = errors.New("empty user")
)

func ValidateCreateOrderParams(r *lomsService.CreateOrderParams) error {
	if r.User == 0 {
		return ErrEmptyUser
	}

	return nil
}

func (impl *implementation) CreateOrder(ctx context.Context, params *lomsService.CreateOrderParams) (*lomsService.CreateOrderResponse, error) {
	if err := ValidateCreateOrderParams(params); err != nil {
		return nil, err
	}

	items := make([]domain.OrderItem, len(params.Items))

	for idx, item := range params.Items {
		items[idx] = domain.OrderItem{
			Sku:   item.Sku,
			Count: item.Count,
		}
	}

	orderId, err := impl.lomsDomain.CreateOrder(params.User, items)

	if err != nil {
		return nil, errors.Wrap(err, "creation failed")
	}

	return &lomsService.CreateOrderResponse{OrderId: orderId}, nil
}
