package service

import (
	"api-repository/internal/config"
	"api-repository/internal/services/file-service/service/internal/handlers"
	"api-repository/pkg/db/minio"
	"context"
	"io"
	"log"
	"net"
	"strconv"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "api-repository/pkg/api/file-service"
)

type FileService struct {
	grpcServer   *grpc.Server
	minioClient  *minio.Client
	sourceClient pb.FileClient
	wg           sync.WaitGroup
	shutdownChan chan struct{}
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

	conn, err := grpc.Dial(
		cfg.UserServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	sourceClient := pb.NewFileClient(conn)
	fileHandler := handlers.NewFileHandler(minioClient)

	grpcServer := grpc.NewServer()
	pb.RegisterFileServer(grpcServer, &fileServer{
		fileHandler: fileHandler,
	})

	return &FileService{
		grpcServer:   grpcServer,
		minioClient:  minioClient,
		sourceClient: sourceClient,
		shutdownChan: make(chan struct{}),
	}, nil
}

func (s *FileService) Start(port int) error {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		return err
	}

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		if err := s.grpcServer.Serve(lis); err != nil {
			log.Printf("gRPC server error: %v", err)
		}
	}()

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.processRequests()
	}()

	return nil
}

func (s *FileService) Stop() {
	close(s.shutdownChan)
	s.grpcServer.GracefulStop()
	s.wg.Wait()
}

func (s *FileService) processRequests() {
	for {
		select {
		case <-s.shutdownChan:
			return
		default:
			stream, err := s.sourceClient.GetFileStream(context.Background(), &pb.EmptyRequest{})
			if err != nil {
				log.Printf("Stream error: %v", err)
				continue
			}
			for {
				req, err := stream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Printf("Receive error: %v", err)
					break
				}
				data, err := s.minioClient.GetFile(context.Background(), req.BucketName, req.ObjectKey)
				if err != nil {
					log.Printf("File read error: %v", err)
					continue
				}
				_, err = s.sourceClient.ReceiveFile(context.Background(), &pb.ReceiveFileRequest{
					BucketName: req.BucketName,
					ObjectKey:  req.ObjectKey,
					Content:    data,
				})
				if err != nil {
					log.Printf("Send file error: %v", err)
				}
			}
		}
	}
}
