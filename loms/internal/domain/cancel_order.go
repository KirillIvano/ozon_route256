package domain

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
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

	if status == "ORDER_CANCELLED" || status == "ORDER_COMPLETED" {
		return ErrCancelOrderWrongStatus
	}

	err = m.lomsRepository.RunTransaction(ctx, pgx.ReadCommitted, func(tx context.Context) error {
		if err := m.lomsRepository.UpdateOrderStatus(tx, orderId, "ORDER_CANCELLED"); err != nil {
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
