package checkout_server_v1

import (
	"context"
	checkoutV1 "route256/checkout/pkg/checkout_v1"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	ErrAddToCartEmptyUser  = errors.New("empty user")
	ErrAddToCartEmptySKU   = errors.New("empty sku")
	ErrAddToCartBadRequest = errors.New("bad request")
)

func ValidateAddToCart(r *checkoutV1.AddToCartParams) error {
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

func (impl *implementation) AddToCart(ctx context.Context, req *checkoutV1.AddToCartParams) (*emptypb.Empty, error) {
	if err := ValidateAddToCart(req); err != nil {
		return nil, err
	}

	err := impl.checkoutDomain.AddToCart(ctx, req.User, req.Sku, req.Count)

	if err != nil {
		return nil, errors.Wrap(err, "creation failed")
	}

	return &emptypb.Empty{}, nil
}
