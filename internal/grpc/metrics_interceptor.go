package grpc

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

var (
	requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_requests_total",
			Help: "Total number of gRPC requests received.",
		},
		[]string{"method"},
	)

	requestsLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "grpc_requests_latency_seconds",
			Help:    "Latency of gRPC requests in seconds.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method"},
	)

	requestsErrorsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_requests_errors_total",
			Help: "Total number of gRPC errors.",
		},
		[]string{"method", "code"},
	)

	concurrentRequests = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "grpc_concurrent_requests",
			Help: "Current number of concurrent gRPC requests.",
		},
		[]string{"method"},
	)
)

func init() {

	prometheus.MustRegister(requestsTotal)
	prometheus.MustRegister(requestsLatency)
	prometheus.MustRegister(requestsErrorsTotal)
	prometheus.MustRegister(concurrentRequests)
}

func MetricsInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		method := info.FullMethod

		requestsTotal.WithLabelValues(method).Inc()

		concurrentRequests.WithLabelValues(method).Inc()

		startTime := time.Now()
		resp, err := handler(ctx, req)
		latency := time.Since(startTime).Seconds()

		requestsLatency.WithLabelValues(method).Observe(latency)

		concurrentRequests.WithLabelValues(method).Dec()

		if err != nil {
			code := status.Code(err).String()
			requestsErrorsTotal.WithLabelValues(method, code).Inc()
		}

		return resp, err
	}
}
