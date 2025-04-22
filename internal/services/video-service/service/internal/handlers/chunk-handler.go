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
)

type ChunkHandler struct {
	minioConn *minio.Client
}

func NewChunkHandler(m *minio.Client) *ChunkHandler {
	return &ChunkHandler{
		minioConn: m,
	}
}

const bucketName = "videos"

func (vh *ChunkHandler) GetVideo(ctx context.Context, req *videoService.GetVideoChunkRequest) (*videoService.GetVideoChunkResponse, error) {
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

func (vh *ChunkHandler) SetVideo(ctx context.Context, req *videoService.SetVideoChunkRequest) (*videoService.SetVideoChunkResponse, error) {

	objName := uuid.New()
	videoData := req.ChunkData

	reader := bytes.NewReader(videoData)
	_, err := vh.minioConn.AppendObject(ctx, bucketName, objName.String(), reader, int64(len(videoData)), minio.AppendObjectOptions{
		Progress:             nil,
		ChunkSize:            1 >> 20,
		DisableContentSha256: false,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &videoService.SetVideoChunkResponse{
		VideoId: objName.String(),
	}, nil
}
