package domain

import (
	"context"

	"github.com/pkg/errors"
)

func (m *CheckoutDomain) Purchase(ctx context.Context, userId int64) (int64, error) {
	items, err := m.repository.GetCartItems(ctx, int64(userId))
	if err != nil {
		return 0, errors.Wrap(err, "getting items from database")
	}
	orderId, err := m.lomsService.CreateOrder(ctx, userId, items)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create order")
	}

	// наверное, лучше унести в кафку, пока не обрабатываю
	go func() {
		m.repository.DeleteCart(ctx, userId)
	}()

	return orderId, err
}
