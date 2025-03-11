package handlers

import (
	"context"

	pb "github.com/PlegunovN/testTask"
)

func (h *FileHandler) UploadFile(ctx context.Context, req *pb.UploadRequest) (*pb.UploadResponse, error) {

	return h.svc.UploadFile(ctx, req)
}
