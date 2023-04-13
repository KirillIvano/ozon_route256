package domain

import (
	"context"
)

func (m *LomsDomain) ListOrder(ctx context.Context, orderId int64) (*OrderInfo, error) {
	order, err := m.lomsRepository.GetOrderInfo(ctx, orderId)

	if err != nil {
		return nil, err
	}

	return order, nil
}
