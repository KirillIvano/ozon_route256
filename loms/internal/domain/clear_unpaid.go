package domain

import (
	"context"

	"github.com/pkg/errors"
)

func (m LomsDomain) ClearUnpaid(ctx context.Context) error {
	orders, err := m.lomsRepository.GetOutdatedOrders(ctx)

	if err != nil {
		return errors.Wrap(err, "reading outdated orders")
	}

	for _, orderId := range orders {
		err := m.lomsRepository.RunReadCommitedTransaction(ctx, func(txCtx context.Context) error {
			err := m.lomsRepository.ClearReservations(txCtx, orderId)
			if err != nil {
				return errors.Wrap(err, "clearing reservation")
			}

			err = m.lomsRepository.UpdateOrderStatus(txCtx, orderId, OrderStatusFailed)
			if err != nil {
				return errors.Wrap(err, "updating order status")
			}

			err = m.orderSender.SendOrder(txCtx, orderId, OrderStatusFailed)
			if err != nil {
				return errors.Wrap(err, "failed to send new order status")
			}

			return nil
		})

		if err != nil {
			return errors.Wrap(err, "update unpaid transaction failed")
		}
	}

	return nil
}
