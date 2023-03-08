package loms_server_v1

import (
	"context"
	lomsV1 "route256/loms/pkg/loms_v1"

	"github.com/pkg/errors"
)

var (
	ErrListOrderEmptyOrderId = errors.New("empty order id")
)

func ValidateListOrderParams(r *lomsV1.ListOrderParams) error {
	if r.OrderID == 0 {
		return ErrListOrderEmptyOrderId
	}

	return nil
}

func (impl *implementation) ListOrder(ctx context.Context, params *lomsV1.ListOrderParams) (*lomsV1.ListOrderResponse, error) {
	if err := ValidateListOrderParams(params); err != nil {
		return nil, err
	}

	orderInfo, err := impl.lomsDomain.ListOrder(params.OrderID)

	if err != nil {
		return nil, errors.Wrap(err, "creation failed")
	}

	responseItems := make([]*lomsV1.OrderItem, len(orderInfo.Items))
	for idx, item := range orderInfo.Items {
		responseItems[idx] = &lomsV1.OrderItem{
			Sku:   item.Sku,
			Count: item.Count,
		}
	}

	return &lomsV1.ListOrderResponse{
		Status: lomsV1.OrderStatus(lomsV1.OrderStatus_value[orderInfo.Status]),
		User:   orderInfo.User,
		Items:  responseItems,
	}, nil
}
