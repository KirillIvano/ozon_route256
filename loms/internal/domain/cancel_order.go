package domain

import (
	"context"
	"errors"
	"log"
)

var (
	ErrCancelOrderUpdateFailed = errors.New("cancel order update failed")
	ErrCancelOrderWrongStatus  = errors.New("wrong order status")
)

func (m *LomsDomain) CancelOrder(ctx context.Context, orderId int64) error {
	status, err := m.lomsRepository.GetOrderStatus(ctx, orderId)

	if err != nil {
		return err
	}

	if status == OrderStatusCancelled || status == OrderStatusPayed {
		return ErrCancelOrderWrongStatus
	}

	err = m.lomsRepository.RunReadCommitedTransaction(ctx, func(tx context.Context) error {
		if err := m.lomsRepository.UpdateOrderStatus(tx, orderId, OrderStatusPayed); err != nil {
			log.Println(err)
			return ErrCancelOrderUpdateFailed
		}

		if err := m.lomsRepository.ClearReservations(tx, orderId); err != nil {
			log.Println(err)
			return ErrCancelOrderUpdateFailed
		}

		return nil
	})

	return err
}
