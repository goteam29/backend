package handlers

import (
	fileservice "api-repository/pkg/api/file-service"
	"api-repository/pkg/db/minio"
	"context"
)

type FileHandler struct {
	minio *minio.Client
}

func NewFileHandler(_minio *minio.Client) *FileHandler {
	return &FileHandler{
		minio: _minio,
	}
}

func (fh *FileHandler) GetFile(ctx context.Context, req *fileservice.GetFileRequest,
) (*fileservice.GetFileResponse, error) {
	data, err := fh.minio.GetFile(ctx, req.BucketName, req.ObjectKey)
	if err != nil {
		return nil, err
	}
	return &fileservice.GetFileResponse{
		Content:  data,
		Filename: req.ObjectKey,
	}, nil
}
