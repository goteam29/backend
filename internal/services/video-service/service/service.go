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

func (vs *VideoService) GetVideo(ctx context.Context, req *videoService.GetVideoChunkRequest) (*videoService.GetVideoChunkResponse, error) {
	return vs.videoHandler.GetVideo(ctx, req)
}

func (vs *VideoService) SetVideo(ctx context.Context, req *videoService.SetVideoChunkRequest) (*videoService.SetVideoChunkResponse, error) {
	return vs.videoHandler.SetVideo(ctx, req)
}
