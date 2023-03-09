package domain

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

func (m *CheckoutDomain) Purchase(ctx context.Context, userId int64) error {
	// TODO temporary variable, delete when database appears
	itemsMock := []CartItem{
		{
			UserId: 2,
			Sku:    2,
			Count:  3,
		},
	}

	items := itemsMock
	orderId, err := m.lomsService.CreateOrder(ctx, userId, items)

	if err != nil {
		return errors.Wrap(err, "failed to create order")
	}
	fmt.Printf("order with id %d has been created\n", orderId)

	return nil
}
