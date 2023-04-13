package universal_server

import (
	"context"
	"route256/libs/logger"
	"strconv"
	"time"

	"net/http"
	"runtime"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	namespace = "route256"
)

type serverMetrics struct {
	requestCounter  *prometheus.CounterVec
	goroutinesGauge *prometheus.Gauge
	responseTime    *prometheus.HistogramVec
}

type MetricsManager struct {
	serverMetrics serverMetrics
}

func createBaseMetrics(project string) serverMetrics {
	requestCounter := promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: project,
		Name:      "requests_total",
	}, []string{"method"})

	goroutinesGauge := promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: project,
		Name:      "goroutines_total",
	})

	responseTime := promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: project,
		Name:      "response_time",
		Buckets:   prometheus.ExponentialBuckets(0.0001, 2, 16),
	}, []string{"status", "method"})

	return serverMetrics{
		requestCounter,
		&goroutinesGauge,
		responseTime,
	}
}

func (m MetricsManager) startObserver(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(time.Second)

		for {
			select {
			case <-ticker.C:
				(*m.serverMetrics.goroutinesGauge).Add(float64(runtime.NumGoroutine()))
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (m MetricsManager) metricsInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	m.serverMetrics.requestCounter.WithLabelValues(info.FullMethod).Inc()
	timeStart := time.Now()

	res, err := handler(ctx, req)

	elapsed := time.Since(timeStart)

	code, isStatus := status.FromError(err)

	if !isStatus {
		m.serverMetrics.responseTime.WithLabelValues("unknown", info.FullMethod).Observe(elapsed.Seconds())
	} else {
		stringStatus := strconv.Itoa(int(code.Code()))
		m.serverMetrics.responseTime.WithLabelValues(stringStatus, info.FullMethod).Observe(elapsed.Seconds())
	}

	return res, err
}

func (m MetricsManager) startMetricsServer(ctx context.Context, address string) {
	srv := http.Server{Addr: address}

	m.startObserver(ctx)

	http.Handle("/metrics", promhttp.Handler())

	go func() {
		err := srv.ListenAndServe()

		if err != nil {
			logger.Fatal("metric server failed to start", zap.Error(err))
		}
	}()
}

func newMetricsManager(project string) *MetricsManager {
	baseMetrics := createBaseMetrics(project)

	return &MetricsManager{
		baseMetrics,
	}
}
