package handlers

import (
	fileservice "api-repository/pkg/api/file-service"
	"api-repository/pkg/utils"
	"bytes"
	"context"
	"errors"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

type FileHandler struct {
	Minio *minio.Client
}

func NewFileHandler(_minio *minio.Client) *FileHandler {
	return &FileHandler{
		Minio: _minio,
	}
}

func (fh *FileHandler) GetFile(ctx context.Context, req *fileservice.GetFileRequest) (*fileservice.GetFileResponse, error) {
	info, err := fh.Minio.StatObject(ctx, req.BucketName, req.ObjectKey, minio.StatObjectOptions{})
	if err != nil {
		var minioErr minio.ErrorResponse
		if errors.As(err, &minioErr) && minioErr.Code == "NoSuchKey" {
			return nil, status.Error(codes.NotFound, "not found")
		}
		return nil, status.Error(codes.Internal, "ISE")
	}

	file, err := fh.Minio.GetObject(ctx, req.BucketName, req.ObjectKey, minio.GetObjectOptions{})
	if err != nil {
		return nil, status.Error(codes.Internal, "ISE")
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, status.Error(codes.Internal, "ISE")
	}

	return &fileservice.GetFileResponse{
		Content:      data,
		Filename:     req.ObjectKey,
		ContentType:  info.ContentType,
		Size:         info.Size,
		LastModified: info.LastModified.String(),
	}, nil
}

func (fh *FileHandler) SetFile(ctx context.Context, req *fileservice.SetFileRequest) (*fileservice.SetFileResponse, error) {

	reader := bytes.NewReader(req.Object)
	_, err := fh.Minio.PutObject(ctx, req.BucketName, req.ObjectName, reader, int64(len(req.Object)),
		minio.PutObjectOptions{
			ContentType: "application/octet-stream",
		})
	if err != nil {
		utils.GetSugaredLogger().Logf(zapcore.ErrorLevel, "error uploading file | err: %v", err)
		return nil, status.Error(codes.Internal, "error uploading file")
	}
	return &fileservice.SetFileResponse{}, nil
}
