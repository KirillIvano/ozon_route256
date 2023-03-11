package checkout_server

import (
	"context"
	checkoutService "route256/checkout/pkg/checkout_service"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	ErrDeleteFromCartEmptyUser  = errors.New("empty user")
	ErrDeleteFromCartEmptySKU   = errors.New("empty sku")
	ErrDeleteFromCartBadRequest = errors.New("bad request")
)

func ValidateDeleteFromCartRequest(r *checkoutService.DeleteFromCartParams) error {
	if r.User == 0 {
		return ErrDeleteFromCartEmptyUser
	}

	if r.Sku == 0 {
		return ErrDeleteFromCartEmptySKU
	}

	if r.Count <= 0 {
		return ErrDeleteFromCartBadRequest
	}

	return nil
}

func (impl *implementation) DeleteFromCart(ctx context.Context, req *checkoutService.DeleteFromCartParams) (*emptypb.Empty, error) {
	if err := ValidateDeleteFromCartRequest(req); err != nil {
		return nil, err
	}

	err := impl.checkoutDomain.DeleteFromCart(ctx, req.User, req.Sku, req.Count)

	if err != nil {
		return nil, errors.Wrap(err, "deletion failed")
	}

	return &emptypb.Empty{}, nil
}
