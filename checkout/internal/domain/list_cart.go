package domain

import "github.com/pkg/errors"

type ListCartResponse struct {
	Offers     []Offer
	TotalPrice uint32
}

func (m *CheckoutDomain) ListCart(user uint32) (ListCartResponse, error) {
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

	var priceChan = make(chan uint32)
	var errChan = make(chan error)

	for idx, item := range items {
		go func(idx int, item CartItem) {
			product, err := m.productService.GetProduct(item.Sku)

			if err != nil {
				errChan <- err
				return
			}

			// разные области памяти, безопасно
			res[idx] = Offer{
				CartItem: item,
				Price:    product.Price,
				Name:     product.Name,
			}
			priceChan <- product.Price
		}(idx, item)
	}

	totalPrice := uint32(0)
	for i := 0; i < len(items); i++ {
		select {
		case price := <-priceChan:
			totalPrice += price
		case err := <-errChan:
			return ListCartResponse{}, errors.Wrap(err, "fetching products")
		}
	}

	return ListCartResponse{
		Offers:     res,
		TotalPrice: totalPrice,
	}, nil
}
