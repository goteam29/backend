package handlers_test

import (
	"api-repository/internal/services/file-service/service/internal/handlers"
	fileservice "api-repository/pkg/api/file-service"
	file_minio "api-repository/pkg/db/file-minio"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockMinioClient struct {
	mock.Mock
}

func (m *MockMinioClient) GetFileWithMeta(ctx context.Context, bucketName, objectKey string) (*file_minio.FileObject, error) {
	args := m.Called(ctx, bucketName, objectKey)
	return args.Get(0).(*file_minio.FileObject), args.Error(1)
}
func TestGetFile_Success(t *testing.T) {
	mockMinio := new(MockMinioClient)
	handler := handlers.NewFileHandler(mockMinio)

	expectedFile := &file_minio.FileObject{
		Content:      []byte("test content"),
		ContentType:  "text/plain",
		Size:         12,
		LastModified: time.Now(),
	}

	req := &fileservice.GetFileRequest{
		BucketName: "test-bucket",
		ObjectKey:  "file.txt",
	}

	mockMinio.On("GetFileWithMeta", mock.Anything, "test-bucket", "file.txt").Return(expectedFile, nil)

	resp, err := handler.GetFile(context.Background(), req)

	assert.NoError(t, err)
	assert.Equal(t, expectedFile.Content, resp.Content)
	assert.Equal(t, expectedFile.ContentType, resp.ContentType)
	assert.Equal(t, expectedFile.Size, resp.Size)
	assert.WithinDuration(t, expectedFile.LastModified, resp.LastModified.AsTime(), time.Second)
	mockMinio.AssertExpectations(t)
}
func TestGetFile_ErrorFromMinio(t *testing.T) {
	mockMinio := new(MockMinioClient)
	handler := handlers.NewFileHandler(mockMinio)

	req := &fileservice.GetFileRequest{
		BucketName: "test-bucket",
		ObjectKey:  "missing.txt",
	}

	mockMinio.On("GetFileWithMeta", mock.Anything, "test-bucket", "missing.txt").Return(nil, errors.New("file not found"))

	resp, err := handler.GetFile(context.Background(), req)

	assert.Nil(t, resp)
	assert.EqualError(t, err, "file not found")
	mockMinio.AssertExpectations(t)
}
