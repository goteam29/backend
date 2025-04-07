package service

import (
	"api-repository/internal/config"
	"api-repository/internal/services/file-service/service/internal/handlers"
	"api-repository/pkg/db/minio"
	"context"
	"net"
	"strconv"

	"google.golang.org/grpc"

	pb "api-repository/pkg/api/file-service"
)

type FileService struct {
	grpcServer  *grpc.Server
	minioClient *minio.Client
}

type fileServer struct {
	pb.UnimplementedFileServer
	fileHandler *handlers.FileHandler
	minioClient *minio.Client
}

func (s *fileServer) GetFile(ctx context.Context, req *pb.GetFileRequest) (*pb.GetFileResponse, error) {
	return s.fileHandler.GetFile(ctx, req)
}

func New(cfg *config.MainConfig) (*FileService, error) {
	minioClient, err := minio.New(cfg.MinIO)
	if err != nil {
		return nil, err
	}

	grpcServer := grpc.NewServer()
	pb.RegisterFileServer(grpcServer, &fileServer{
		minioClient: minioClient,
		fileHandler: handlers.NewFileHandler(minioClient),
	})

	return &FileService{
		grpcServer:  grpcServer,
		minioClient: minioClient,
	}, nil
}

func (s *FileService) Start(port int) error {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return err
	}

	return s.grpcServer.Serve(lis)
}

func (s *FileService) Stop() {
	s.grpcServer.GracefulStop()
}
