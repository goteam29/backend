//go:build test

package service

import (
	"api-repository/internal/services/file-service/service/internal/handlers"
	fileservice "api-repository/pkg/api/file-service"
	file_minio "api-repository/pkg/db/file-minio"
	"context"
)

type FileServerForTest struct {
	fileservice.UnimplementedFileServer
	FileHandler *handlers.FileHandler
	MinioClient *file_minio.Client
}

func (s *FileServerForTest) GetFile(ctx context.Context, req *fileservice.GetFileRequest) (*fileservice.GetFileResponse, error) {
	return s.FileHandler.GetFile(ctx, req)
}
