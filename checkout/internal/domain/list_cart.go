package domain

import (
	"github.com/pkg/errors"
)

type ListCartResponse struct {
	Offers     []Offer
	TotalPrice uint32
}

func (m *Model) ListCart(user uint32) (ListCartResponse, error) {
	itemsMock := []CartItem{
		{
			UserId: 2,
			Sku:    1148162,
			Count:  3,
		},
		{
			UserId: 2,
			Sku:    1625903,
			Count:  3,
		},
	}
	items := itemsMock

	res := make([]Offer, len(items))
	totalPrice := uint32(0)

	// TODO: ваще такое параллелить надо, но пока у меня лапки
	for idx, item := range items {
		product, err := m.productService.GetProduct(item.Sku)

		if err != nil {
			return ListCartResponse{}, errors.Wrap(err, "requesting product info")
		}

		res[idx] = Offer{
			CartItem: item,
			Price:    product.Price,
			Name:     product.Name,
		}
		totalPrice += product.Price
	}

	return ListCartResponse{
		Offers:     res,
		TotalPrice: totalPrice,
	}, nil
}
