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

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
)

func prepareAddToCartTest(
	t *testing.T,
	ctx context.Context,
	getStocks func(ctx context.Context, sku uint32) ([]domain.Stock, error),
	addToCart func(ctx context.Context, userId int64, sku uint32, count uint32) error,
) *domain.CheckoutDomain {
	controller := minimock.NewController(t)

	productsMock := mProduct.NewProductsClientMock(controller)
	lomsMock := mLoms.NewLomsClientMock(controller)
	repository := mRepo.NewCheckoutRepositoryMock(controller)
	wpMock := worker_pool.New(ctx, 10)

	lomsMock.StocksMock.Set(getStocks)
	repository.AddToCartMock.Set(addToCart)

	model := domain.New(lomsMock, productsMock, repository, wpMock)

	return model
}

func TestAddToCart(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		model := prepareAddToCartTest(
			t,
			ctx,
			func(ctx context.Context, sku uint32) ([]domain.Stock, error) {
				return []domain.Stock{
					{
						WarehouseID: 1,
						Count:       10,
					},
					{
						WarehouseID: 2,
						Count:       10,
					},
				}, nil
			},
			func(ctx context.Context, userId int64, sku uint32, count uint32) error {
				return nil
			},
		)

		err := model.AddToCart(ctx, 10, 10, 20)

		require.Nil(t, err)
	})

	t.Run("insufficient returns error", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		model := prepareAddToCartTest(
			t,
			ctx,
			func(ctx context.Context, sku uint32) ([]domain.Stock, error) {
				return []domain.Stock{
					{
						WarehouseID: 1,
						Count:       5,
					},
				}, nil
			},
			func(ctx context.Context, userId int64, sku uint32, count uint32) error {
				return nil
			},
		)
		err := model.AddToCart(ctx, 10, 10, 10)

		require.ErrorIs(t, err, domain.ErrInsufficientStocks)
	})

	t.Run("stocks returns error", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		testErr := errors.New("test error")

		model := prepareAddToCartTest(
			t,
			ctx,
			func(ctx context.Context, sku uint32) ([]domain.Stock, error) {
				return nil, testErr
			},
			func(ctx context.Context, userId int64, sku uint32, count uint32) error {
				return nil
			},
		)
		err := model.AddToCart(ctx, 10, 10, 10)

		require.ErrorIs(t, err, testErr)
	})

	t.Run("insufficient stocks", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()
		testErr := errors.New("test error")

		model := prepareAddToCartTest(
			t,
			ctx,
			func(ctx context.Context, sku uint32) ([]domain.Stock, error) {
				return []domain.Stock{
					{
						WarehouseID: 1,
						Count:       10,
					},
				}, nil
			},
			func(ctx context.Context, userId int64, sku uint32, count uint32) error {
				return testErr
			},
		)
		err := model.AddToCart(ctx, 10, 10, 10)

		require.ErrorIs(t, err, testErr)
	})
}
