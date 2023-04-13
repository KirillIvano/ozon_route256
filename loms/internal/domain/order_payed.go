package domain

import (
	"context"
	"route256/libs/logger"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var (
	ErrOrderPayedWrongStatus = errors.New("wrong order status")
)

func (m *LomsDomain) SetOrderPayedTransaction(ctx context.Context, orderId int64) error {
	if err := m.lomsRepository.ApplyOrderReservations(ctx, orderId); err != nil {
		return errors.Wrap(err, "reservations apply failed")
	}
	logger.Info("payment: applied reservations", zap.Int64("orderId", orderId))

	if err := m.lomsRepository.ClearReservations(ctx, orderId); err != nil {
		return errors.Wrap(err, "reservations clear failed")
	}
	logger.Info("payment: cleared reservations", zap.Int64("orderId", orderId))

	if err := m.lomsRepository.UpdateOrderStatus(ctx, orderId, OrderStatusPayed); err != nil {
		return errors.Wrap(err, "status update failed")
	}
	logger.Info("payment: updated status", zap.Int64("orderId", orderId))

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
