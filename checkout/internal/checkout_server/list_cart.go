package checkout_server

import (
	"context"
	checkoutService "route256/checkout/pkg/checkout_service"

	"github.com/pkg/errors"
)

var (
	ErrListCartEmptyUser = errors.New("empty user")
)

func ValidateListCart(r *checkoutService.ListCartParams) error {
	if r.User == 0 {
		return ErrListCartEmptyUser
	}

	return nil
}

func (impl *implementation) ListCart(ctx context.Context, req *checkoutService.ListCartParams) (*checkoutService.ListCartResponse, error) {
	if err := ValidateListCart(req); err != nil {
		return nil, err
	}

	res, err := impl.checkoutDomain.ListCart(uint32(req.User))

	if err != nil {
		return nil, errors.Wrap(err, "getting failed")
	}

	items := make([]*checkoutService.ListCartItem, len(res.Offers))
	for i, offer := range res.Offers {
		items[i] = &checkoutService.ListCartItem{
			Sku:   offer.Sku,
			Count: offer.Count,
			Name:  offer.Name,
			Price: uint32(offer.Price),
		}
	}

	return &checkoutService.ListCartResponse{
		Items:      items,
		TotalPrice: res.TotalPrice,
	}, nil
}
