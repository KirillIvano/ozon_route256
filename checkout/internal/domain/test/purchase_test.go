package domain_test

import (
	"context"
	"errors"
	mLoms "route256/checkout/internal/clients/loms/mocks"
	mProduct "route256/checkout/internal/clients/products/mocks"
	"route256/checkout/internal/domain"
	mRepo "route256/checkout/internal/repository/mocks"
	"route256/checkout/pkg/worker_pool"
	"testing"
	"time"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func prepareTest(
	t *testing.T,
	ctx context.Context,
	getCartItems func(ctx context.Context, userId int64) (ca1 []domain.CartItem, err error),
	createOrder func(ctx context.Context, user int64, items []domain.CartItem) (int64, error),
	deleteCart func(ctx context.Context, userId int64) (err error),
) *domain.CheckoutDomain {
	controller := minimock.NewController(t)

	productsMock := mProduct.NewProductsClientMock(controller)
	lomsMock := mLoms.NewLomsClientMock(controller)
	repository := mRepo.NewCheckoutRepositoryMock(controller)
	wpMock := worker_pool.New(ctx, 10)

	repository.GetCartItemsMock.Set(getCartItems)
	lomsMock.CreateOrderMock.Set(createOrder)
	repository.DeleteCartMock.Set(deleteCart)

	model := domain.New(lomsMock, productsMock, repository, wpMock)

	return model
}

func TestPurchase(t *testing.T) {

	t.Run("happy path", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		controller := minimock.NewController(t)

		productsMock := mProduct.NewProductsClientMock(controller)
		lomsMock := mLoms.NewLomsClientMock(controller)
		repository := mRepo.NewCheckoutRepositoryMock(controller)
		wpMock := worker_pool.New(ctx, 10)

		repository.GetCartItemsMock.Set(func(ctx context.Context, userId int64) (ca1 []domain.CartItem, err error) {
			return []domain.CartItem{}, nil
		})
		lomsMock.CreateOrderMock.Set(func(ctx context.Context, user int64, items []domain.CartItem) (int64, error) {
			return 10, nil
		})
		repository.DeleteCartMock.When(ctx, 10).Then(nil)

		model := domain.New(lomsMock, productsMock, repository, wpMock)

		res, err := model.Purchase(ctx, 10)

		controller.Wait(time.Millisecond)

		require.True(t, repository.MinimockDeleteCartDone())
		require.Nil(t, err)
		require.Equal(t, int64(10), res)
	})

	t.Run("getting cart items returns error", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		testErr := errors.New("test error")

		model := prepareTest(
			t,
			ctx,
			func(ctx context.Context, userId int64) (ca1 []domain.CartItem, err error) {
				return nil, testErr
			},
			func(ctx context.Context, user int64, items []domain.CartItem) (int64, error) {
				return 10, nil
			},
			func(ctx context.Context, userId int64) (err error) { return nil },
		)
		_, err := model.Purchase(ctx, 10)

		require.ErrorIs(t, err, testErr)
	})

	t.Run("creating order returns error", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		testErr := errors.New("test error")

		model := prepareTest(
			t,
			ctx,
			func(ctx context.Context, userId int64) (ca1 []domain.CartItem, err error) {
				return []domain.CartItem{}, nil
			},
			func(ctx context.Context, user int64, items []domain.CartItem) (int64, error) {
				return 0, testErr
			},
			func(ctx context.Context, userId int64) (err error) { return nil },
		)
		_, err := model.Purchase(ctx, 10)

		require.ErrorIs(t, err, testErr)
	})
}
