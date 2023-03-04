package loms_server_v1

import (
	"context"
	"route256/loms/internal/domain"
	lomsV1 "route256/loms/pkg/loms_v1"

	"github.com/pkg/errors"
)

var (
	ErrEmptyUser = errors.New("empty user")
)

func ValidateCreateOrderParams(r *lomsV1.CreateOrderParams) error {
	if r.User == 0 {
		return ErrEmptyUser
	}

	return nil
}

func (impl *implementation) CreateOrder(ctx context.Context, params *lomsV1.CreateOrderParams) (*lomsV1.CreateOrderResponse, error) {
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

	return &lomsV1.CreateOrderResponse{OrderId: orderId}, nil
}
