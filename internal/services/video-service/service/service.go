package service

import (
	"api-repository/internal/services/video-service/service/internal/handlers"
	videoService "api-repository/pkg/api/video-service"
	"context"
	"github.com/minio/minio-go/v7"
)

type VideoService struct {
	videoService.UnimplementedVideoServer
	minioConn    *minio.Client
	videoHandler *handlers.ChunkHandler
}

func NewVideoService(m *minio.Client) *VideoService {
	return &VideoService{
		minioConn:    m,
		videoHandler: handlers.NewChunkHandler(m),
	}
}

func (vs *VideoService) GetVideoChunk(ctx context.Context, req *videoService.GetVideoChunkRequest) (*videoService.GetVideoChunkResponse, error) {
	return vs.videoHandler.GetVideoChunk(ctx, req)
}

func (vs *VideoService) SetVideoChunk(ctx context.Context, req *videoService.SetVideoChunkRequest) (*videoService.SetVideoChunkResponse, error) {
	return vs.videoHandler.SetVideoChunk(ctx, req)
}

func (vs *VideoService) AddToVideoChunk(ctx context.Context, req *videoService.AddToVideoChunkRequest) (*videoService.AddToVideoChunkResponse, error) {
	return vs.videoHandler.AddToVideoChunk(ctx, req)
}
