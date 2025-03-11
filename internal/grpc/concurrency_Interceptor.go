package grpc

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func ConcurrencyInterceptor(uploadDownloadSem, listSem chan struct{}) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		var sem chan struct{}
		switch info.FullMethod {
		case "/pb.FileService/UploadFile", "/pb.FileService/DownloadFile":
			sem = uploadDownloadSem
		case "/pb.FileService/ListFiles":
			sem = listSem
		default:
			return handler(ctx, req)
		}

		select {
		case sem <- struct{}{}:
			defer func() { <-sem }()
			return handler(ctx, req)
		default:
			return nil, fmt.Errorf("too many simultaneous requests")
		}
	}
}
