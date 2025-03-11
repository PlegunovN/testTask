package handlers

import (
	"context"

	pb "github.com/PlegunovN/testTask"
)

func (h *FileHandler) DownloadFile(ctx context.Context, req *pb.DownloadRequest) (*pb.DownloadResponse, error) {

	return h.svc.DownloadFile(ctx, req)
}
