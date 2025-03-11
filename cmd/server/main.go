package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/PlegunovN/testTask"
	"github.com/PlegunovN/testTask/internal/configs"
	fileservice "github.com/PlegunovN/testTask/internal/files"
	grpcinterceptor "github.com/PlegunovN/testTask/internal/grpc"
	"github.com/PlegunovN/testTask/internal/handlers"
	"github.com/PlegunovN/testTask/internal/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {

	logger := logger.InitLogger()
	defer func() {
		if err := logger.Sync(); err != nil {
			fmt.Printf("Error syncing logger: %v\n", err)
		}
	}()

	cfg, err := configs.LoadConfig(".env")
	if err != nil {
		logger.Fatal("error load config", zap.Error(err))
	}

	port := flag.Int("port", cfg.PORTGRPC, "gRPC server port")
	metricsPort := flag.Int("metrics_port", cfg.PORTMETRICS, "Port for exposing metrics")
	flag.Parse()

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		logger.Info("Metrics are available on the port %d in /metrics", zap.Int("port", cfg.PORTMETRICS))
		if err := http.ListenAndServe(fmt.Sprintf(":%d", *metricsPort), nil); err != nil {
			logger.Fatal("HTTP server error for metrics: %v", zap.Error(err))
		}
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logger.Fatal("Failed to open port: %v", zap.Error(err))
	}

	var (
		uploadDownloadSem = make(chan struct{}, cfg.UPLOADLIMIT)
		listSem           = make(chan struct{}, cfg.LISTLIMIT)
	)

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			grpcinterceptor.MetricsInterceptor(),
			grpcinterceptor.ValidationInterceptor(),
			grpcinterceptor.ConcurrencyInterceptor(uploadDownloadSem, listSem),
		),
	}
	server := grpc.NewServer(opts...)

	svc, err := fileservice.NewService(cfg.STORAGEDIR, logger)
	if err != nil {
		logger.Fatal("Failed to initialize file service: %v", zap.Error(err))
	}

	handler := handlers.NewFileHandler(svc)
	pb.RegisterFileServiceServer(server, handler)

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		logger.Info("Shutting down the server...")
		server.GracefulStop()
	}()

	logger.Info("gRPC server running on port %d", zap.Int("port", cfg.PORTGRPC))
	if err := server.Serve(lis); err != nil {
		logger.Fatal("Server error: %v", zap.Error(err))
	}
}
