package loms_server_v1

import (
	"context"
	lomsV1 "route256/loms/pkg/loms_v1"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	ErrCancelOrderEmptyOrderId = errors.New("empty order id")
)

func ValidateCancelOrderParams(r *lomsV1.OrderCancelParams) error {
	if r.OrderID == 0 {
		return ErrCancelOrderEmptyOrderId
	}

	return nil
}

func (impl *implementation) CancelOrder(ctx context.Context, params *lomsV1.OrderCancelParams) (*emptypb.Empty, error) {
	if err := ValidateCancelOrderParams(params); err != nil {
		return nil, err
	}

	err := impl.lomsDomain.CancelOrder(params.OrderID)

	if err != nil {
		return nil, errors.Wrap(err, "cancellation failed")
	}

	return &emptypb.Empty{}, nil
}
