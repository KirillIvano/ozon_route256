package domain

import (
	"errors"
)

var (
	ErrorCreateOrderInvalidItems = errors.New("invalid items count")
)

func (m *LomsDomain) CreateOrder(user int64, items []OrderItem) (int64, error) {
	if len(items) == 0 {
		return 0, ErrorCreateOrderInvalidItems
	}

	return 123, nil
}
