package loms_server

import (
	"context"
	lomsService "route256/loms/pkg/loms_service"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	ErrCancelOrderEmptyOrderId = errors.New("empty order id")
)

func ValidateCancelOrderParams(r *lomsService.OrderCancelParams) error {
	if r.OrderID == 0 {
		return ErrCancelOrderEmptyOrderId
	}

	return nil
}

func (impl *implementation) CancelOrder(ctx context.Context, params *lomsService.OrderCancelParams) (*emptypb.Empty, error) {
	if err := ValidateCancelOrderParams(params); err != nil {
		return nil, err
	}

	err := impl.lomsDomain.CancelOrder(ctx, params.OrderID)

	if err != nil {
		return nil, errors.Wrap(err, "cancellation failed")
	}

	return &emptypb.Empty{}, nil
}
