package domain

import (
	"context"
	"math"

	"github.com/pkg/errors"
)

var (
	ErrorCreateOrderInvalidItems   = errors.New("invalid items count")
	ErrorReservationNotEnoughItems = errors.New("not enough items")
)

func CalculateItemReservation(item OrderItem, stocks []Stock, orderId int64) ([]Reservation, error) {
	reservations := make([]Reservation, 0)
	currentCnt := int32(item.Count)

	for _, stock := range stocks {
		diff := int32(math.Min(float64(currentCnt), float64(stock.Count)))
		if diff == 0 {
			continue
		}

		reservations = append(reservations, Reservation{
			Sku:         item.Sku,
			Count:       diff,
			WarehouseId: stock.WarehouseID,
			OrderId:     orderId,
		})

		currentCnt -= diff
		if currentCnt == 0 {
			break
		}
	}

	if currentCnt != 0 {
		return nil, ErrorReservationNotEnoughItems
	}

	return reservations, nil
}

func (m LomsDomain) CreateOrderReservations(ctx context.Context, orderId int64, items []OrderItem) error {
	reservations := make([]Reservation, 0, len(items))

	for _, item := range items {
		stocks, err := m.lomsRepository.ListStocks(ctx, item.Sku)
		if err != nil {
			return err
		}

		itemReservations, err := CalculateItemReservation(item, stocks, orderId)
		if err != nil {
			return err
		}

		reservations = append(reservations, itemReservations...)
	}

	err := m.lomsRepository.CreateReservation(ctx, reservations)

	return err
}

func (m LomsDomain) CreateOrderItemsTransaction(ctx context.Context, orderId int64, items []OrderItem) error {
	err := m.lomsRepository.CreateOrderItems(ctx, orderId, items)
	if err != nil {
		return errors.Wrap(err, "creating order items")
	}

	err = m.CreateOrderReservations(ctx, orderId, items)
	if err != nil {
		return errors.Wrap(err, "creating order reservations")
	}

	return nil
}

func (m LomsDomain) CreateOrder(ctx context.Context, user int64, items []OrderItem) (int64, error) {
	if len(items) == 0 {
		return 0, ErrorCreateOrderInvalidItems
	}

	orderId, err := m.lomsRepository.CreateOrder(ctx, user)
	if err != nil {
		return 0, errors.Wrap(err, "creating order")
	}
	err = m.orderSender.SendOrder(ctx, orderId, OrderStatusNew)
	if err != nil {
		return 0, errors.Wrap(err, "failed to send new order status")
	}

	// Сохраняем элементы заказа и резервируем под них места на складах
	err = m.lomsRepository.RunReadCommitedTransaction(ctx, func(ctxTX context.Context) error {
		return m.CreateOrderItemsTransaction(ctx, orderId, items)
	})

	// В зависимости от результата резервации выставляем статус
	if err != nil {
		if err := m.lomsRepository.UpdateOrderStatus(ctx, orderId, OrderStatusFailed); err != nil {
			return 0, errors.Wrap(err, "updating failed order status")
		}

		err = m.orderSender.SendOrder(ctx, orderId, OrderStatusFailed)
		if err != nil {
			return 0, errors.Wrap(err, "failed to send new order status")
		}
	} else {
		if err := m.lomsRepository.UpdateOrderStatus(ctx, orderId, OrderStatusAwaitingPayment); err != nil {
			return 0, errors.Wrap(err, "updating successful order status")
		}

		err = m.orderSender.SendOrder(ctx, orderId, OrderStatusAwaitingPayment)
		if err != nil {
			return 0, errors.Wrap(err, "failed to send new order status")
		}
	}

	return orderId, nil
}
