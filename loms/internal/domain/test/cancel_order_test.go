package domain_test

import (
	"context"
	"errors"
	"route256/loms/internal/domain"
	"route256/loms/internal/order_sender/mocks"
	mRepo "route256/loms/internal/repository/mocks"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func TestCancelOrder(t *testing.T) {
	testError := errors.New("test error")

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		controller := minimock.NewController(t)

		sender := mocks.NewOrderSenderMock(controller)
		repository := mRepo.NewLomsRepositoryMock(controller)
		model := domain.New(repository, sender)

		sender.SendOrderMock.Set(func(ctx context.Context, orderId int64, orderStatus string) (err error) { return nil })
		repository.GetOrderStatusMock.When(ctx, 10).Then(domain.OrderStatusAwaitingPayment, nil)
		repository.RunReadCommitedTransactionMock.Set(func(c context.Context, fx func(ctxTX context.Context) error) (err error) {
			return nil
		})

		err := model.CancelOrder(ctx, 10)

		require.Nil(t, err)
	})

	t.Run("status getting failed", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		controller := minimock.NewController(t)

		sender := mocks.NewOrderSenderMock(controller)
		repository := mRepo.NewLomsRepositoryMock(controller)
		model := domain.New(repository, sender)

		sender.SendOrderMock.Set(func(ctx context.Context, orderId int64, orderStatus string) (err error) { return nil })
		repository.GetOrderStatusMock.When(ctx, 10).Then("", testError)

		err := model.CancelOrder(ctx, 10)

		require.ErrorIs(t, err, testError)
	})

	t.Run("invalid status", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		controller := minimock.NewController(t)

		sender := mocks.NewOrderSenderMock(controller)
		repository := mRepo.NewLomsRepositoryMock(controller)
		model := domain.New(repository, sender)

		sender.SendOrderMock.Set(func(ctx context.Context, orderId int64, orderStatus string) (err error) { return nil })
		repository.GetOrderStatusMock.When(ctx, 10).Then(domain.OrderStatusCancelled, nil)

		err := model.CancelOrder(ctx, 10)

		require.ErrorIs(t, err, domain.ErrCancelOrderWrongStatus)
	})

	t.Run("failed transaction", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		controller := minimock.NewController(t)

		sender := mocks.NewOrderSenderMock(controller)
		repository := mRepo.NewLomsRepositoryMock(controller)
		model := domain.New(repository, sender)

		sender.SendOrderMock.Set(func(ctx context.Context, orderId int64, orderStatus string) (err error) { return nil })
		repository.GetOrderStatusMock.When(ctx, 10).Then(domain.OrderStatusAwaitingPayment, nil)
		repository.RunReadCommitedTransactionMock.Set(func(c context.Context, fx func(ctxTX context.Context) error) (err error) {
			return testError
		})

		err := model.CancelOrder(ctx, 10)

		require.ErrorIs(t, err, testError)
	})

	t.Run("status update failed", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		controller := minimock.NewController(t)

		sender := mocks.NewOrderSenderMock(controller)
		repository := mRepo.NewLomsRepositoryMock(controller)
		model := domain.New(repository, sender)

		sender.SendOrderMock.Set(func(ctx context.Context, orderId int64, orderStatus string) (err error) { return nil })
		repository.UpdateOrderStatusMock.When(ctx, 100, domain.OrderStatusPayed).Then(testError)
		err := model.ClearOrderInfoTransaction(ctx, 100)

		require.ErrorIs(t, err, domain.ErrCancelOrderUpdateFailed)
	})

	t.Run("clear reservations failed", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		controller := minimock.NewController(t)

		sender := mocks.NewOrderSenderMock(controller)
		repository := mRepo.NewLomsRepositoryMock(controller)
		model := domain.New(repository, sender)

		sender.SendOrderMock.Set(func(ctx context.Context, orderId int64, orderStatus string) (err error) { return nil })
		repository.UpdateOrderStatusMock.When(ctx, 100, domain.OrderStatusPayed).Then(nil)
		repository.ClearReservationsMock.When(ctx, 100).Then(testError)

		err := model.ClearOrderInfoTransaction(ctx, 100)

		require.ErrorIs(t, err, domain.ErrCancelOrderUpdateFailed)
	})

	t.Run("success transaction", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		controller := minimock.NewController(t)

		sender := mocks.NewOrderSenderMock(controller)
		repository := mRepo.NewLomsRepositoryMock(controller)
		model := domain.New(repository, sender)

		sender.SendOrderMock.Set(func(ctx context.Context, orderId int64, orderStatus string) (err error) { return nil })
		repository.UpdateOrderStatusMock.When(ctx, 100, domain.OrderStatusPayed).Then(nil)
		repository.ClearReservationsMock.When(ctx, 100).Then(nil)

		err := model.ClearOrderInfoTransaction(ctx, 100)

		require.Nil(t, err)
	})
}
