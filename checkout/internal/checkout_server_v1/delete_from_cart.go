package checkout_server_v1

import (
	"context"
	checkoutV1 "route256/checkout/pkg/checkout_v1"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	ErrDeleteFromCartEmptyUser  = errors.New("empty user")
	ErrDeleteFromCartEmptySKU   = errors.New("empty sku")
	ErrDeleteFromCartBadRequest = errors.New("bad request")
)

func ValidateDeleteFromCartRequest(r *checkoutV1.DeleteFromCartParams) error {
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

func (impl *implementation) DeleteFromCart(ctx context.Context, req *checkoutV1.DeleteFromCartParams) (*emptypb.Empty, error) {
	if err := ValidateDeleteFromCartRequest(req); err != nil {
		return nil, err
	}

	err := impl.checkoutDomain.DeleteFromCart(req.User, req.Sku, req.Count)

	if err != nil {
		return nil, errors.Wrap(err, "deletion failed")
	}

	return &emptypb.Empty{}, nil
}
