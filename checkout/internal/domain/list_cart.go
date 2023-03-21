package domain

import (
	"context"

	"github.com/pkg/errors"
)

type ListCartResponse struct {
	Offers     []Offer
	TotalPrice uint32
}

// func EmulateRequest(ctx context.Context, someId uint32) (ProductInfo, error) {
// 	select {
// 	case <-time.After(2 * time.Second):
// 		return ProductInfo{}, errors.New("timeout happened")
// 	case <-ctx.Done():
// 		return ProductInfo{}, errors.New("context fucked up")
// 	}
// }

func (m *CheckoutDomain) ListCart(ctx context.Context, userId uint32) (*ListCartResponse, error) {
	items, err := m.repository.GetCartItems(ctx, int64(userId))
	if err != nil {
		return nil, errors.Wrap(err, "getting items from database")
	}

	res := make([]Offer, len(items))

	priceChan := make(chan uint32)
	defer close(priceChan)

	errChan := make(chan error)
	defer close(errChan)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for idx, item := range items {
		currentIdx, currentItem := idx, item

		m.wp.Run(func() {
			product, err := m.productService.GetProduct(ctx, currentItem.Sku)

			if err != nil {
				errChan <- err
				return
			}

			res[currentIdx] = Offer{
				CartItem: item,
				Price:    product.Price,
				Name:     product.Name,
			}
			priceChan <- product.Price
		})
	}

	totalPrice := uint32(0)
	var firstError error

	for i := 0; i < len(items); i++ {
		select {
		case price := <-priceChan:
			totalPrice += price
		case err := <-errChan:
			if firstError == nil {
				firstError = err
				// сообщаем контексту, что остальные запросы не нужны
				cancel()
			}
		}
	}

	if firstError != nil {
		return nil, errors.Wrap(firstError, "getting info from products service")
	}

	return &ListCartResponse{
		Offers:     res,
		TotalPrice: totalPrice,
	}, nil
}
