package domain

import (
	"context"

	"github.com/pkg/errors"
)

func (m *LomsDomain) Stocks(ctx context.Context, sku uint32) ([]Stock, error) {
	res, err := m.lomsRepository.ListStocks(ctx, sku)

	if err != nil {
		return nil, errors.Wrap(err, "accessing stocks database")
	}

	return res, nil
}
