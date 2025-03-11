package handlers

import (
	"context"

	pb "github.com/PlegunovN/testTask"
)

func (h *FileHandler) ListFiles(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {

	return h.svc.ListFiles(ctx, req)
}
