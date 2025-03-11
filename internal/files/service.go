package files

import (
	"context"
	"os"
	"path/filepath"
	"sync"
	"time"

	pb "github.com/PlegunovN/testTask"
	"go.uber.org/zap"
)

type Service struct {
	storageDir string
	mu         sync.Mutex
	metadata   map[string]*FileMeta
	logger     *zap.Logger
}

type FileMeta struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewService(storageDir string, logger *zap.Logger) (*Service, error) {

	err := os.MkdirAll(storageDir, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return &Service{
		storageDir: storageDir,
		metadata:   make(map[string]*FileMeta),
		logger:     logger,
	}, nil
}

func (s *Service) UploadFile(ctx context.Context, req *pb.UploadRequest) (*pb.UploadResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	filePath := filepath.Join(s.storageDir, req.Filename)
	now := time.Now()

	meta, exists := s.metadata[req.Filename]
	if !exists {
		meta = &FileMeta{CreatedAt: now}
	}
	meta.UpdatedAt = now

	err := os.WriteFile(filePath, req.Data, 0644)
	if err != nil {
		s.logger.Info("file writing error: %v", zap.Error(err))
		return nil, err
	}

	s.metadata[req.Filename] = meta

	return &pb.UploadResponse{Message: "file uploaded successfully"}, nil
}

func (s *Service) DownloadFile(ctx context.Context, req *pb.DownloadRequest) (*pb.DownloadResponse, error) {
	filePath := filepath.Join(s.storageDir, req.Filename)
	data, err := os.ReadFile(filePath)
	if err != nil {
		s.logger.Info("file read error: %v", zap.Error(err))
		return nil, err
	}
	return &pb.DownloadResponse{Data: data}, nil
}

func (s *Service) ListFiles(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var files []*pb.FileInfo
	for filename, meta := range s.metadata {
		files = append(files, &pb.FileInfo{
			Filename:  filename,
			CreatedAt: meta.CreatedAt.Format(time.RFC3339),
			UpdatedAt: meta.UpdatedAt.Format(time.RFC3339),
		})
	}
	return &pb.ListResponse{Files: files}, nil
}
