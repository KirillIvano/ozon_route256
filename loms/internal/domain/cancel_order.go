package domain

import (
	"context"

	"github.com/pkg/errors"
)

var (
	ErrCancelOrderUpdateFailed = errors.New("cancel order update failed")
	ErrCancelOrderWrongStatus  = errors.New("wrong order status")
)

func (m *LomsDomain) ClearOrderInfoTransaction(tx context.Context, orderId int64) error {
	if err := m.lomsRepository.UpdateOrderStatus(tx, orderId, OrderStatusCancelled); err != nil {
		return ErrCancelOrderUpdateFailed
	}

	if err := m.lomsRepository.ClearReservations(tx, orderId); err != nil {
		return ErrCancelOrderUpdateFailed
	}

	err := m.orderSender.SendOrder(tx, orderId, OrderStatusCancelled)
	if err != nil {
		return errors.Wrap(err, "failed to send new order status")
	}

	return nil
}

func (m *LomsDomain) CancelOrder(ctx context.Context, orderId int64) error {
	status, err := m.lomsRepository.GetOrderStatus(ctx, orderId)

	if err != nil {
		return err
	}

	if status == OrderStatusCancelled || status == OrderStatusPayed {
		return ErrCancelOrderWrongStatus
	}

	err = m.lomsRepository.RunReadCommitedTransaction(ctx, func(ctxTX context.Context) error {
		return m.ClearOrderInfoTransaction(ctx, orderId)
	})

	return err
}
