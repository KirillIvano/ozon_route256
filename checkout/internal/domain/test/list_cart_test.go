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

func prepareListCart(
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

func TestListCart(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		controller := minimock.NewController(t)

		productsMock := mProduct.NewProductsClientMock(controller)
		lomsMock := mLoms.NewLomsClientMock(controller)
		repository := mRepo.NewCheckoutRepositoryMock(controller)
		wpMock := worker_pool.New(ctx, 10)

		cartItems := []domain.CartItem{
			{
				Sku:    1,
				UserId: 1,
				Count:  10,
			},
			{
				Sku:    2,
				UserId: 2,
				Count:  10,
			},
		}

		firstProduct := domain.ProductInfo{
			Name:  "first",
			Price: 100,
		}
		secondProduct := domain.ProductInfo{
			Name:  "second",
			Price: 100,
		}

		repository.GetCartItemsMock.When(ctx, 10).Then(cartItems, nil)
		productsMock.GetProductMock.Set(func(_ context.Context, sku uint32) (domain.ProductInfo, error) {
			if sku == 1 {
				return firstProduct, nil
			}
			if sku == 2 {
				return secondProduct, nil
			}

			return domain.ProductInfo{}, errors.New("")
		})

		model := domain.New(lomsMock, productsMock, repository, wpMock)
		res, err := model.ListCart(ctx, 10)

		require.Nil(t, err)

		require.Equal(t, uint32(2000), res.TotalPrice)

		for _, offer := range res.Offers {
			switch offer.Name {
			case "first":
				require.Equal(t, domain.Offer{
					CartItem: cartItems[0],
					Price:    100,
					Name:     "first",
				}, offer)
			case "second":
				require.Equal(t, domain.Offer{
					CartItem: cartItems[1],
					Price:    100,
					Name:     "second",
				}, offer)
			default:
				require.FailNow(t, "wrong name")
			}
		}
	})

	t.Run("error in product getting", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		controller := minimock.NewController(t)

		productsMock := mProduct.NewProductsClientMock(controller)
		lomsMock := mLoms.NewLomsClientMock(controller)
		repository := mRepo.NewCheckoutRepositoryMock(controller)
		wpMock := worker_pool.New(ctx, 10)

		testError := errors.New("test error")

		repository.GetCartItemsMock.When(ctx, 10).Then([]domain.CartItem{
			{
				Sku:    1,
				UserId: 1,
				Count:  10,
			},
		}, nil)
		productsMock.GetProductMock.Set(func(_ context.Context, sku uint32) (domain.ProductInfo, error) {
			return domain.ProductInfo{}, testError
		})

		model := domain.New(lomsMock, productsMock, repository, wpMock)
		_, err := model.ListCart(ctx, 10)

		require.ErrorIs(t, err, testError)
	})

	t.Run("error in getting cart items", func(t *testing.T) {
		t.Parallel()

		ctx := context.Background()

		controller := minimock.NewController(t)

		productsMock := mProduct.NewProductsClientMock(controller)
		lomsMock := mLoms.NewLomsClientMock(controller)
		repository := mRepo.NewCheckoutRepositoryMock(controller)
		wpMock := worker_pool.New(ctx, 10)

		testError := errors.New("test error")

		repository.GetCartItemsMock.When(ctx, 10).Then([]domain.CartItem{}, testError)
		productsMock.GetProductMock.Set(func(_ context.Context, sku uint32) (domain.ProductInfo, error) {
			return domain.ProductInfo{}, nil
		})

		model := domain.New(lomsMock, productsMock, repository, wpMock)
		_, err := model.ListCart(ctx, 10)

		require.ErrorIs(t, err, testError)
	})
}
