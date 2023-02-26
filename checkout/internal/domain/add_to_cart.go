package domain

import "github.com/pkg/errors"

var (
	ErrInsufficientStocks = errors.New("insufficient stocks")
)

func (m *CheckoutDomain) AddToCart(user int64, sku uint32, count uint16) error {
	stocks, err := m.lomsService.Stocks(sku)
	if err != nil {
		return errors.WithMessage(err, "checking stocks")
	}

	counter := int64(count)
	for _, stock := range stocks {
		counter -= int64(stock.Count)
		if counter <= 0 {
			return nil
		}
	}

	return ErrInsufficientStocks
}
