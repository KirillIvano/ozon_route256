package checkout_server_v1

import (
	"context"
	checkoutV1 "route256/checkout/pkg/checkout_v1"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	ErrPurchaseEmptyUser = errors.New("empty user")
)

func ValidatePurchaseRequest(r *checkoutV1.PurchaseParams) error {
	if r.User == 0 {
		return ErrPurchaseEmptyUser
	}

	return nil
}

func (impl *implementation) Purchase(ctx context.Context, req *checkoutV1.PurchaseParams) (*emptypb.Empty, error) {
	if err := ValidatePurchaseRequest(req); err != nil {
		return nil, err
	}

	err := impl.checkoutDomain.Purchase(ctx, req.User)

	if err != nil {
		return nil, errors.Wrap(err, "purchase failed")
	}

	return &emptypb.Empty{}, nil
}
