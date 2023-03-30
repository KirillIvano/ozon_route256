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

func TestListOrder(t *testing.T) {
	testError := errors.New("test error")

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		controller := minimock.NewController(t)

		repository := mRepo.NewLomsRepositoryMock(controller)
		model := domain.New(repository)

		repository.GetOrderInfoMock.When(ctx, 10).Then(&domain.OrderInfo{
			Status: domain.OrderStatusAwaitingPayment,
			User:   10,
			Items:  []domain.OrderItem{},
		}, nil)

		_, err := model.ListOrder(ctx, 10)

		require.Nil(t, err)
	})

	t.Run("error listing", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		controller := minimock.NewController(t)

		repository := mRepo.NewLomsRepositoryMock(controller)
		model := domain.New(repository)

		repository.GetOrderInfoMock.When(ctx, 10).Then(&domain.OrderInfo{}, testError)

		_, err := model.ListOrder(ctx, 10)

		require.ErrorIs(t, err, testError)
	})
}
