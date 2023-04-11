package metrics

import (
	"context"
	"net/http"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	namespace = "route256"
	subsystem = "grpc"
)

var (
	reqCounter = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "requests_total",
	}, []string{"method"})

	goroutinesGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "goroutines_total",
	})

	histogramResponseTime = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "response_time",
		Buckets:   prometheus.ExponentialBuckets(0.0001, 2, 16),
	}, []string{"status", "method"})
)

type Metrics struct{}

func (m *Metrics) Observe(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(time.Second)

		for {
			select {
			case <-ticker.C:
				goroutinesGauge.Add(float64(runtime.NumGoroutine()))
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (m *Metrics) Handler() http.Handler {
	return promhttp.Handler()
}

func New() *Metrics {
	return &Metrics{}
}
