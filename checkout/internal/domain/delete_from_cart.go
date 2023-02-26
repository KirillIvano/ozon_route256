package domain

import (
	"github.com/pkg/errors"
)

var (
	ErrItemDontExist = errors.New("item doesn't exist")
	ErrInvalidCount  = errors.New("count is invalid")
)

func (m *CheckoutDomain) DeleteFromCart(user int64, sku uint32, count uint16) error {
	itemMock := CartItem{
		UserId: 2,
		Sku:    2,
		Count:  3,
	}

	item := itemMock
	// if err != nil {
	// 	return ErrItemDontExist
	// }

	if item.Count < count {
		return ErrInvalidCount
	}

	// do deletion

	return nil
}
