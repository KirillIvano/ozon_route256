package checkout_server

import (
	"context"
	checkoutService "route256/checkout/pkg/checkout_service"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	ErrAddToCartEmptyUser  = errors.New("empty user")
	ErrAddToCartEmptySKU   = errors.New("empty sku")
	ErrAddToCartBadRequest = errors.New("bad request")
)

func ValidateAddToCart(r *checkoutService.AddToCartParams) error {
	if r.User == 0 {
		return ErrAddToCartEmptyUser
	}

	if r.Sku == 0 {
		return ErrAddToCartEmptySKU
	}

	if r.Count <= 0 {
		return ErrAddToCartBadRequest
	}

	return nil
}

func (impl *implementation) AddToCart(ctx context.Context, req *checkoutService.AddToCartParams) (*emptypb.Empty, error) {
	if err := ValidateAddToCart(req); err != nil {
		return nil, err
	}

	err := impl.checkoutDomain.AddToCart(ctx, req.User, req.Sku, req.Count)

	if err != nil {
		return nil, errors.Wrap(err, "add to cart failed")
	}

	return &emptypb.Empty{}, nil
}
