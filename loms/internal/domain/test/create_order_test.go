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

func TestCreateOrder(t *testing.T) {
	orderItems := []domain.OrderItem{
		{
			Sku:   1,
			Count: 10,
		},
	}
	testError := errors.New("test error")

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		controller := minimock.NewController(t)

		sender := mocks.NewOrderSenderMock(controller)
		repository := mRepo.NewLomsRepositoryMock(controller)
		model := domain.New(repository, sender)

		sender.SendOrderMock.Set(func(ctx context.Context, orderId int64, orderStatus string) (err error) { return nil })
		repository.CreateOrderMock.When(ctx, 10).Then(100, nil)
		repository.CreateOrderItemsMock.When(ctx, 100, orderItems).Then(nil)
		repository.ListStocksMock.Set(func(ctx context.Context, sku uint32) (sa1 []domain.Stock, err error) {
			return []domain.Stock{
				{
					WarehouseID: 1,
					Count:       10,
				},
			}, nil
		})
		repository.CreateReservationMock.Set(func(ctx context.Context, reservations []domain.Reservation) (err error) {
			return nil
		})
		repository.UpdateOrderStatusMock.Set(func(ctx context.Context, orderId int64, status string) (err error) {
			return nil
		})
		repository.RunReadCommitedTransactionMock.Set(func(c context.Context, fx func(ctxTX context.Context) error) (err error) {
			return fx(ctx)
		})

		res, err := model.CreateOrder(ctx, 10, orderItems)

		require.Nil(t, err)
		require.Equal(t, int64(100), res)
	})

	t.Run("create order failed", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		controller := minimock.NewController(t)

		sender := mocks.NewOrderSenderMock(controller)
		repository := mRepo.NewLomsRepositoryMock(controller)
		model := domain.New(repository, sender)

		sender.SendOrderMock.Set(func(ctx context.Context, orderId int64, orderStatus string) (err error) { return nil })
		repository.CreateOrderMock.When(ctx, 10).Then(0, testError)

		_, err := model.CreateOrder(ctx, 10, orderItems)

		require.ErrorIs(t, err, testError)
	})

	t.Run("create order items failed", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		controller := minimock.NewController(t)

		sender := mocks.NewOrderSenderMock(controller)
		repository := mRepo.NewLomsRepositoryMock(controller)
		model := domain.New(repository, sender)

		sender.SendOrderMock.Set(func(ctx context.Context, orderId int64, orderStatus string) (err error) { return nil })
		repository.CreateOrderMock.When(ctx, 10).Then(100, nil)
		repository.CreateOrderItemsMock.When(ctx, 100, orderItems).Then(testError)

		err := model.CreateOrderItemsTransaction(ctx, 100, orderItems)

		require.ErrorIs(t, err, testError)
	})

	t.Run("list stocks failed", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		controller := minimock.NewController(t)

		sender := mocks.NewOrderSenderMock(controller)
		repository := mRepo.NewLomsRepositoryMock(controller)
		model := domain.New(repository, sender)

		sender.SendOrderMock.Set(func(ctx context.Context, orderId int64, orderStatus string) (err error) { return nil })
		repository.CreateOrderMock.When(ctx, 10).Then(100, nil)
		repository.CreateOrderItemsMock.When(ctx, 100, orderItems).Then(nil)
		repository.ListStocksMock.Set(func(ctx context.Context, sku uint32) (sa1 []domain.Stock, err error) {
			return nil, testError
		})

		err := model.CreateOrderItemsTransaction(ctx, 100, orderItems)

		require.ErrorIs(t, err, testError)
	})

	t.Run("not enough items", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		controller := minimock.NewController(t)

		sender := mocks.NewOrderSenderMock(controller)
		repository := mRepo.NewLomsRepositoryMock(controller)
		model := domain.New(repository, sender)

		sender.SendOrderMock.Set(func(ctx context.Context, orderId int64, orderStatus string) (err error) { return nil })
		repository.CreateOrderMock.When(ctx, 10).Then(100, nil)
		repository.CreateOrderItemsMock.When(ctx, 100, orderItems).Then(nil)
		repository.ListStocksMock.Set(func(ctx context.Context, sku uint32) (sa1 []domain.Stock, err error) {
			return []domain.Stock{
				{
					WarehouseID: 1,
					Count:       0,
				},
				{
					WarehouseID: 1,
					Count:       5,
				},
			}, nil
		})

		err := model.CreateOrderItemsTransaction(ctx, 100, orderItems)

		require.ErrorIs(t, err, domain.ErrorReservationNotEnoughItems)
	})

	t.Run("success order status error", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		controller := minimock.NewController(t)

		sender := mocks.NewOrderSenderMock(controller)
		repository := mRepo.NewLomsRepositoryMock(controller)
		model := domain.New(repository, sender)

		sender.SendOrderMock.Set(func(ctx context.Context, orderId int64, orderStatus string) (err error) { return nil })
		repository.CreateOrderMock.When(ctx, 10).Then(100, nil)
		repository.RunReadCommitedTransactionMock.Set(func(c context.Context, fx func(ctxTX context.Context) error) (err error) {
			return nil
		})
		repository.UpdateOrderStatusMock.Set(func(ctx context.Context, orderId int64, status string) (err error) {
			return testError
		})

		_, err := model.CreateOrder(ctx, 10, orderItems)

		require.ErrorIs(t, err, testError)
	})

	t.Run("failed order status error", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		controller := minimock.NewController(t)

		sender := mocks.NewOrderSenderMock(controller)
		repository := mRepo.NewLomsRepositoryMock(controller)
		model := domain.New(repository, sender)

		sender.SendOrderMock.Set(func(ctx context.Context, orderId int64, orderStatus string) (err error) { return nil })
		repository.CreateOrderMock.When(ctx, 10).Then(100, nil)
		repository.RunReadCommitedTransactionMock.Set(func(c context.Context, fx func(ctxTX context.Context) error) (err error) {
			return testError
		})
		repository.UpdateOrderStatusMock.Set(func(ctx context.Context, orderId int64, status string) (err error) {
			return testError
		})

		_, err := model.CreateOrder(ctx, 10, orderItems)

		require.ErrorIs(t, err, testError)
	})

	t.Run("validation failed", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		controller := minimock.NewController(t)
		sender := mocks.NewOrderSenderMock(controller)
		repository := mRepo.NewLomsRepositoryMock(controller)
		model := domain.New(repository, sender)

		sender.SendOrderMock.Set(func(ctx context.Context, orderId int64, orderStatus string) (err error) { return nil })

		_, err := model.CreateOrder(ctx, 10, []domain.OrderItem{})

		require.ErrorIs(t, err, domain.ErrorCreateOrderInvalidItems)
	})
}
