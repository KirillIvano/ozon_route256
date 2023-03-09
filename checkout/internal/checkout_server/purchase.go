package checkout_server

import (
	"context"
	checkoutService "route256/checkout/pkg/checkout_service"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	ErrPurchaseEmptyUser = errors.New("empty user")
)

func ValidatePurchaseRequest(r *checkoutService.PurchaseParams) error {
	if r.User == 0 {
		return ErrPurchaseEmptyUser
	}

	return nil
}

func (impl *implementation) Purchase(ctx context.Context, req *checkoutService.PurchaseParams) (*emptypb.Empty, error) {
	if err := ValidatePurchaseRequest(req); err != nil {
		return nil, err
	}

	err := impl.checkoutDomain.Purchase(ctx, req.User)

	if err != nil {
		return nil, errors.Wrap(err, "purchase failed")
	}

	return &emptypb.Empty{}, nil
}
