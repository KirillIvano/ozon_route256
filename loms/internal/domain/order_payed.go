package domain

import (
	"context"
	"log"

	"github.com/pkg/errors"
)

var (
	ErrOrderPayedWrongStatus = errors.New("wrong order status")
)

func (m *LomsDomain) SetOrderPayedTransaction(ctx context.Context, orderId int64) error {
	if err := m.lomsRepository.ApplyOrderReservations(ctx, orderId); err != nil {
		return errors.Wrap(err, "reservations apply failed")
	}
	log.Println("payment: applied reservations")

	if err := m.lomsRepository.ClearReservations(ctx, orderId); err != nil {
		return errors.Wrap(err, "reservations clear failed")
	}
	log.Println("payment: cleared reservations")

	if err := m.lomsRepository.UpdateOrderStatus(ctx, orderId, OrderStatusPayed); err != nil {
		return errors.Wrap(err, "status update failed")
	}
	log.Println("payment: updated status")

	return nil
}

func (m *LomsDomain) SetOrderPayed(ctx context.Context, orderId int64) error {
	status, err := m.lomsRepository.GetOrderStatus(ctx, orderId)

	if err != nil {
		return err
	}

	if status != OrderStatusAwaitingPayment {
		return ErrOrderPayedWrongStatus
	}

	err = m.lomsRepository.RunReadCommitedTransaction(ctx, func(ctxTX context.Context) error {
		return m.SetOrderPayedTransaction(ctxTX, orderId)
	})

	err = m.orderSender.SendOrder(ctx, orderId, OrderStatusPayed)
	if err != nil {
		return errors.Wrap(err, "failed to send new order status")
	}

	if err != nil {
		errors.Wrap(err, "pay order transaction failed")
	}

	return nil
}
