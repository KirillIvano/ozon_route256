package checkout_server_v1

import (
	"context"
	checkoutV1 "route256/checkout/pkg/checkout_v1"

	"github.com/pkg/errors"
)

var (
	ErrListCartEmptyUser = errors.New("empty user")
)

func ValidateListCart(r *checkoutV1.ListCartParams) error {
	if r.User == 0 {
		return ErrListCartEmptyUser
	}

	return nil
}

func (impl *implementation) ListCart(ctx context.Context, req *checkoutV1.ListCartParams) (*checkoutV1.ListCartResponse, error) {
	if err := ValidateListCart(req); err != nil {
		return nil, err
	}

	res, err := impl.checkoutDomain.ListCart(uint32(req.User))

	if err != nil {
		return nil, errors.Wrap(err, "getting failed")
	}

	items := make([]*checkoutV1.ListCartItem, len(res.Offers))
	for i, offer := range res.Offers {
		items[i] = &checkoutV1.ListCartItem{
			Sku:   offer.Sku,
			Count: offer.Count,
			Name:  offer.Name,
			Price: uint32(offer.Price),
		}
	}

	return &checkoutV1.ListCartResponse{
		Items:      items,
		TotalPrice: res.TotalPrice,
	}, nil
}
