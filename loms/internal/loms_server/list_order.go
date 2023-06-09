package loms_server

import (
	"context"
	lomsService "route256/loms/pkg/loms_service"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrListOrderEmptyOrderId = errors.New("empty order id")
)

func ValidateListOrderParams(r *lomsService.ListOrderParams) error {
	if r.OrderID == 0 {
		return ErrListOrderEmptyOrderId
	}

	return nil
}

func (impl *implementation) ListOrder(ctx context.Context, params *lomsService.ListOrderParams) (*lomsService.ListOrderResponse, error) {
	if err := ValidateListOrderParams(params); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	orderInfo, err := impl.lomsDomain.ListOrder(ctx, params.OrderID)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	responseItems := make([]*lomsService.OrderItem, len(orderInfo.Items))
	for idx, item := range orderInfo.Items {
		responseItems[idx] = &lomsService.OrderItem{
			Sku:   item.Sku,
			Count: item.Count,
		}
	}

	return &lomsService.ListOrderResponse{
		Status: lomsService.OrderStatus(lomsService.OrderStatus_value[orderInfo.Status]),
		User:   orderInfo.User,
		Items:  responseItems,
	}, nil
}
