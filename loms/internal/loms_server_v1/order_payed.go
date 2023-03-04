package loms_server_v1

import (
	"context"
	lomsV1 "route256/loms/pkg/loms_v1"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	ErrOrderPayedEmptyOrderId = errors.New("empty order id")
)

func (impl *implementation) OrderPayed(ctx context.Context, req *lomsV1.OrderPayedParams) (*emptypb.Empty, error) {
	err := impl.lomsDomain.SetOrderPayed(req.OrderID)

	if err != nil {
		return nil, errors.Wrap(err, "cancellation failed")
	}

	return &emptypb.Empty{}, nil
}
