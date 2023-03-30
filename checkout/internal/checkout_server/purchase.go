package checkout_server

import (
	"context"
	checkoutService "route256/checkout/pkg/checkout_service"

	"github.com/pkg/errors"
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

func (impl *implementation) Purchase(ctx context.Context, req *checkoutService.PurchaseParams) (*checkoutService.PurchaseResponse, error) {
	if err := ValidatePurchaseRequest(req); err != nil {
		return nil, err
	}

	orderId, err := impl.checkoutDomain.Purchase(ctx, req.User)

	if err != nil {
		return nil, errors.Wrap(err, "purchase failed")
	}

	return &checkoutService.PurchaseResponse{OrderId: orderId}, nil
}
