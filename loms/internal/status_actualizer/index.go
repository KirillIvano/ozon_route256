package status_actualizer

import (
	"context"
	"route256/libs/logger"
	"time"

	"go.uber.org/zap"
)

type LomsDomain interface {
	ClearUnpaid(ctx context.Context) error
}

type OrderStatusActualizer struct {
	domain LomsDomain
	ctx    context.Context
	ticker *time.Ticker
}

// актуализирует заказы раз в минуту
func (e *OrderStatusActualizer) Start() {
	go func() {
		for {
			select {
			case <-e.ctx.Done():
				return
			case <-e.ticker.C:
				err := e.domain.ClearUnpaid(e.ctx)
				if err != nil {
					logger.Error("unpaid order failed to actualize", zap.Error(err))
					break
				}

				logger.Info("payments actualized")
			}
		}
	}()
}

func (e *OrderStatusActualizer) Close() {
	e.ticker.Stop()
}

func New(ctx context.Context, domain LomsDomain) OrderStatusActualizer {
	return OrderStatusActualizer{
		ticker: time.NewTicker(time.Minute),
		ctx:    ctx,
		domain: domain,
	}
}
