package domain

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
)

func (m *CheckoutDomain) Purchase(ctx context.Context, userId int64) error {
	items, err := m.repository.GetCartItems(ctx, int64(userId))
	if err != nil {
		return errors.Wrap(err, "getting items from database")
	}
	orderId, err := m.lomsService.CreateOrder(ctx, userId, items)
	if err != nil {
		return errors.Wrap(err, "failed to create order")
	}

	// Подчищаем корзину, не влияет на основной пайплайн
	// TODO: подумать, оставить ли так или сделать другой механизм, как временный норм
	err = m.repository.DeleteCart(ctx, userId)
	fmt.Println("deletion error", err)

	fmt.Printf("order with id %d has been created\n", orderId)

	return nil
}
