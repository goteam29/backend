package handlers

import (
	videoService "api-repository/pkg/api/video-service"
	"bytes"
	"context"
	"github.com/google/uuid"
	minio "github.com/minio/minio-go/v7"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"sync"
)

type ChunkHandler struct {
	minioConn *minio.Client
	mu        sync.Mutex
	buff      map[string][]byte
}

func NewChunkHandler(m *minio.Client) *ChunkHandler {
	return &ChunkHandler{
		minioConn: m,
		buff:      make(map[string][]byte),
	}
}

const bucketName = "videos"

func (vh *ChunkHandler) GetVideoChunk(ctx context.Context, req *videoService.GetVideoChunkRequest) (*videoService.GetVideoChunkResponse, error) {
	objectName := req.VideoId

	obj, err := vh.minioConn.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer obj.Close()

	content, err := io.ReadAll(obj)
	if err != nil {
		return nil, err
	}

	return &videoService.GetVideoChunkResponse{
		ChunkData: content,
	}, nil
}

func (vh *ChunkHandler) AddToVideoChunk(ctx context.Context, req *videoService.AddToVideoChunkRequest) (*videoService.AddToVideoChunkResponse, error) {
	vh.mu.Lock()
	defer vh.mu.Unlock()
	vh.buff[req.VideoName] = append(vh.buff[req.VideoName], req.ChunkData...)
	return &videoService.AddToVideoChunkResponse{
		VideoName: req.VideoName,
	}, nil
}

func (vh *ChunkHandler) SetVideoChunk(ctx context.Context, req *videoService.SetVideoChunkRequest) (*videoService.SetVideoChunkResponse, error) {
	objName := uuid.New()

	vh.mu.Lock()
	defer vh.mu.Unlock()
	v, exists := vh.buff[req.VideoName]
	if !exists {
		return nil, status.Error(codes.NotFound, "not found")
	}

	reader := bytes.NewReader(v)
	_, err := vh.minioConn.PutObject(ctx, bucketName, objName.String(), reader, int64(len(v)),
		minio.PutObjectOptions{
			ContentType: "application/octet-stream",
		})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	delete(vh.buff, req.VideoName)

	return &videoService.SetVideoChunkResponse{
		VideoId: objName.String(),
	}, nil
}

//func (vh *ChunkHandler) SetVideoChunk(ctx context.Context, req *videoService.SetVideoChunkRequest) (*videoService.SetVideoChunkResponse, error) {
//	objName := uuid.New()
//	videoData := req.ChunkData
//
//	reader := bytes.NewReader(videoData)
//	_, err := vh.minioConn.PutObject(ctx, bucketName, objName.String(), reader, int64(len(videoData)),
//		minio.PutObjectOptions{
//			ContentType: "application/octet-stream",
//		})
//	if err != nil {
//		return nil, status.Error(codes.Internal, err.Error())
//	}
//
//	return &videoService.SetVideoChunkResponse{
//		VideoId: objName.String(),
//	}, nil
//}
