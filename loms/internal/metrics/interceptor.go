package metrics

import (
	"context"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	reqCounter.Inc()

	timeStart := time.Now()

	res, err := handler(ctx, req)

	elapsed := time.Since(timeStart)

	code, isStatus := status.FromError(err)

	if !isStatus {
		histogramResponseTime.WithLabelValues("unknown").Observe(elapsed.Seconds())
	} else {
		stringStatus := strconv.Itoa(int(code.Code()))
		histogramResponseTime.WithLabelValues(stringStatus).Observe(elapsed.Seconds())
	}

	return res, err
}
