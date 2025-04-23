package service

import (
	"api-repository/internal/config"
	"api-repository/internal/services/file-service/service/internal/handlers"
	minio2 "api-repository/pkg/db/minio"
	"context"
	"github.com/minio/minio-go/v7"
	"log"
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
	minioClient := minio2.NewFileMinioConnection(cfg.MinIO)
	fileHandler := handlers.NewFileHandler(minioClient)

	grpcServer := grpc.NewServer()
	pb.RegisterFileServer(grpcServer, &fileServer{
		fileHandler: fileHandler,
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

	if err := s.grpcServer.Serve(lis); err != nil {
		log.Printf("gRPC server error: %v", err)
	}
	return nil
}

func (s *FileService) Stop() {
	s.grpcServer.GracefulStop()
}
