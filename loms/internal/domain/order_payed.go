package domain

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

var (
	ErrOrderPayedWrongStatus = errors.New("wrong order status")
)

func (m *LomsDomain) SetOrderPayed(ctx context.Context, orderId int64) error {
	status, err := m.lomsRepository.GetOrderStatus(ctx, orderId)

	if err != nil {
		return err
	}

	fmt.Println(status)
	if status != "ORDER_AWAITING_PAYMENT" {
		return ErrCancelOrderWrongStatus
	}

	err = m.lomsRepository.RunTransaction(ctx, pgx.ReadCommitted, func(ctxTX context.Context) error {
		if err := m.lomsRepository.ApplyOrderReservations(ctx, orderId); err != nil {
			return errors.Wrap(err, "reservations apply failed")
		}
		log.Println("payment: applied reservations")

		if err := m.lomsRepository.ClearReservations(ctx, orderId); err != nil {
			return errors.Wrap(err, "reservations clear failed")
		}
		log.Println("payment: cleared reservations")

		if err := m.lomsRepository.UpdateOrderStatus(ctx, orderId, "ORDER_PAYED"); err != nil {
			return errors.Wrap(err, "status update failed")
		}
		log.Println("payment: updated status")

		return nil
	})

	return errors.Wrap(err, "pay order transaction failed")
}
