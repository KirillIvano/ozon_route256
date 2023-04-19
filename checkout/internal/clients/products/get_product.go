package products_client

import (
	"context"
	"fmt"
	"route256/checkout/internal/domain"
	productsService "route256/products/pkg/products_service"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	cacheHitsTotal = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "route256",
		Subsystem: "cache",
		Name:      "hits",
	})

	cacheMissTotal = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "route256",
		Subsystem: "cache",
		Name:      "misses",
	})

	histogramResponseTimeCache = promauto.NewHistogram(prometheus.HistogramOpts{
		Namespace: "route256",
		Subsystem: "cache",
		Name:      "response_time_seconds",
		Buckets:   prometheus.ExponentialBuckets(0.000001, 2, 16),
	})
)

func (c *client) getProductFromCache(sku uint32) (domain.ProductInfo, bool) {
	timeStart := time.Now()
	res, found := c.cache.GetFromCache(strconv.Itoa(int(sku)))
	elapsed := time.Since(timeStart)

	histogramResponseTimeCache.Observe(elapsed.Seconds())

	fmt.Println(sku, res, found)

	if !found {
		cacheMissTotal.Inc()
		return domain.ProductInfo{}, false
	} else {
		cacheHitsTotal.Inc()
		return *res, true
	}
}

func (c *client) getProductFromApi(ctx context.Context, sku uint32) (domain.ProductInfo, error) {
	res, err := c.client.GetProduct(ctx, &productsService.GetProductRequest{Sku: sku, Token: c.token})

	product := domain.ProductInfo{
		Name:  res.Name,
		Price: res.Price,
	}
	c.cache.AddToCache(strconv.Itoa(int(sku)), product)

	if err != nil {
		return domain.ProductInfo{}, err
	}

	return product, err
}

func (c *client) GetProduct(ctx context.Context, sku uint32) (domain.ProductInfo, error) {
	fromCache, found := c.getProductFromCache(sku)

	if found {
		return fromCache, nil
	}

	res, err := c.getProductFromApi(ctx, sku)
	if err != nil {
		return domain.ProductInfo{}, err
	}

	return res, nil
}
