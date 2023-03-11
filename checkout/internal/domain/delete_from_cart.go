package domain

import (
	"context"

	"github.com/pkg/errors"
)

func (m *CheckoutDomain) DeleteFromCart(ctx context.Context, user int64, sku uint32, count uint32) error {
	err := m.repository.DeleteItem(ctx, user, sku, count)

	if err != nil {
		return errors.Wrap(err, "deleting items")
	}

	return nil
}
