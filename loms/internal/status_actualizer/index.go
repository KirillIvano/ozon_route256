package status_actualizer

import (
	"context"
	"log"
	"time"
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
					log.Print(err)
					break
				}

				log.Println("payments actualized")
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
