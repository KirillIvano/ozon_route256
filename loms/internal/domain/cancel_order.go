package domain

import (
	"errors"
)

var (
	ErrCancelOrderDontExist = errors.New("order does not exist")
)

func (m *LomsDomain) CancelOrder(orderId int64) error {
	// TODO ходим в базку, ставим статус
	return ErrCancelOrderDontExist
}
