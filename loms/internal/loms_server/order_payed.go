package loms_server

import (
	"context"
	lomsService "route256/loms/pkg/loms_service"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	ErrOrderPayedEmptyOrderId = errors.New("empty order id")
)

func ValidateOrderPayedParams(r *lomsService.OrderPayedParams) error {
	if r.OrderID == 0 {
		return ErrOrderPayedEmptyOrderId
	}

	return nil
}

func (impl *implementation) OrderPayed(ctx context.Context, params *lomsService.OrderPayedParams) (*emptypb.Empty, error) {
	if err := ValidateOrderPayedParams(params); err != nil {
		return nil, err
	}

	err := impl.lomsDomain.SetOrderPayed(params.OrderID)

	if err != nil {
		return nil, errors.Wrap(err, "cancellation failed")
	}

	return &emptypb.Empty{}, nil
}
