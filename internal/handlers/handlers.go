package handlers

import (
	pb "github.com/PlegunovN/testTask"
	"github.com/PlegunovN/testTask/internal/files"
)

type FileHandler struct {
	pb.UnimplementedFileServiceServer
	svc *files.Service
}

func NewFileHandler(svc *files.Service) *FileHandler {
	return &FileHandler{svc: svc}
}
