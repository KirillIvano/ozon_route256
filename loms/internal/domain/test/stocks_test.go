package domain_test

import (
	"context"
	"errors"
	"route256/loms/internal/domain"
	mRepo "route256/loms/internal/repository/mocks"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestStocks(t *testing.T) {
	testError := errors.New("test error")

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		controller := minimock.NewController(t)

		repository := mRepo.NewLomsRepositoryMock(controller)
		model := domain.New(repository)

		stocks := []domain.Stock{
			{
				WarehouseID: 1,
				Count:       10,
			},
		}

		repository.ListStocksMock.When(ctx, 5).Then(stocks, nil)

		res, err := model.Stocks(ctx, 5)

		require.Nil(t, err)
		require.Equal(t, stocks, res)
	})

	t.Run("sad path", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		controller := minimock.NewController(t)

		repository := mRepo.NewLomsRepositoryMock(controller)
		model := domain.New(repository)

		repository.ListStocksMock.When(ctx, 5).Then(nil, testError)

		_, err := model.Stocks(ctx, 5)

		require.ErrorIs(t, err, testError)
	})
}
