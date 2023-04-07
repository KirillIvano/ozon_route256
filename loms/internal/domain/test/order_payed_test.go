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

func TestSetOrderPayed(t *testing.T) {
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

		err := model.SetOrderPayed(ctx, 10)

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

		err := model.SetOrderPayed(ctx, 10)

		require.ErrorIs(t, err, testError)
	})

	t.Run("wrong status", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		controller := minimock.NewController(t)

		sender := mocks.NewOrderSenderMock(controller)
		repository := mRepo.NewLomsRepositoryMock(controller)
		model := domain.New(repository, sender)

		sender.SendOrderMock.Set(func(ctx context.Context, orderId int64, orderStatus string) (err error) { return nil })
		repository.GetOrderStatusMock.When(ctx, 10).Then(domain.OrderStatusCancelled, nil)

		err := model.SetOrderPayed(ctx, 10)

		require.ErrorIs(t, err, domain.ErrOrderPayedWrongStatus)
	})

	t.Run("transaction happy path", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		controller := minimock.NewController(t)

		sender := mocks.NewOrderSenderMock(controller)
		repository := mRepo.NewLomsRepositoryMock(controller)
		model := domain.New(repository, sender)

		sender.SendOrderMock.Set(func(ctx context.Context, orderId int64, orderStatus string) (err error) { return nil })
		repository.ApplyOrderReservationsMock.When(ctx, 10).Then(nil)
		repository.ClearReservationsMock.When(ctx, 10).Then(nil)
		repository.UpdateOrderStatusMock.When(ctx, 10, domain.OrderStatusPayed).Then(nil)

		err := model.SetOrderPayedTransaction(ctx, 10)

		require.Nil(t, err)
	})

	t.Run("apply reservations failed", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		controller := minimock.NewController(t)

		sender := mocks.NewOrderSenderMock(controller)
		repository := mRepo.NewLomsRepositoryMock(controller)
		model := domain.New(repository, sender)

		sender.SendOrderMock.Set(func(ctx context.Context, orderId int64, orderStatus string) (err error) { return nil })
		repository.ApplyOrderReservationsMock.When(ctx, 10).Then(testError)

		err := model.SetOrderPayedTransaction(ctx, 10)

		require.ErrorIs(t, err, testError)
	})

	t.Run("clear reservations failed", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		controller := minimock.NewController(t)

		sender := mocks.NewOrderSenderMock(controller)
		repository := mRepo.NewLomsRepositoryMock(controller)
		model := domain.New(repository, sender)

		sender.SendOrderMock.Set(func(ctx context.Context, orderId int64, orderStatus string) (err error) { return nil })
		repository.ApplyOrderReservationsMock.When(ctx, 10).Then(nil)
		repository.ClearReservationsMock.When(ctx, 10).Then(testError)

		err := model.SetOrderPayedTransaction(ctx, 10)

		require.ErrorIs(t, err, testError)
	})

	t.Run("udpate order status failed", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		controller := minimock.NewController(t)

		sender := mocks.NewOrderSenderMock(controller)
		repository := mRepo.NewLomsRepositoryMock(controller)
		model := domain.New(repository, sender)

		sender.SendOrderMock.Set(func(ctx context.Context, orderId int64, orderStatus string) (err error) { return nil })
		repository.ApplyOrderReservationsMock.When(ctx, 10).Then(nil)
		repository.ClearReservationsMock.When(ctx, 10).Then(nil)
		repository.UpdateOrderStatusMock.When(ctx, 10, domain.OrderStatusPayed).Then(testError)

		err := model.SetOrderPayedTransaction(ctx, 10)

		require.ErrorIs(t, err, testError)
	})
}
