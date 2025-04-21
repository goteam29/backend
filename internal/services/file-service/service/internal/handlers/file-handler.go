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

func (fh *FileHandler) GetFile(ctx context.Context, req *fileservice.GetFileRequest) (*fileservice.GetFileResponse, error) {
	obj, err := fh.minio.GetFileWithMeta(ctx, req.BucketName, req.ObjectKey)
	if err != nil {
		return nil, err
	}

	return &fileservice.GetFileResponse{
		Content:      obj.Content,
		Filename:     req.ObjectKey,
		ContentType:  obj.ContentType,
		Size:         obj.Size,
		LastModified: obj.LastModified,
	}, nil
}
