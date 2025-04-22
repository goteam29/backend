package main

import (
	"api-repository/internal/adapters/interceptors"
	"api-repository/internal/config"
	"api-repository/internal/services"
	"api-repository/internal/services/video-service/service"
	videoService "api-repository/pkg/api/video-service"
	"api-repository/pkg/db/minio"
	"api-repository/pkg/utils"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func main() {
	utils.CreateNewSugaredLogger()

	cfg, err := config.NewMainConfig()
	if err != nil {
		log.Fatalf("can't get config | err: %v", err)
	}

	minioConn := minio.NewVideoMinioConnection(cfg.MinIO)

	svc := service.NewVideoService(minioConn)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.VideServicePort))
	if err != nil {
		utils.GetSugaredLogger().Fatalf("can't start video-service | err %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptors.LoggingInterceptor(utils.GetSugaredLogger())),
	}
	server := grpc.NewServer(opts...)
	videoService.RegisterVideoServer(server, svc)

	log.Printf("Configuration:\n%s", services.GetBeautifulConfigurationString(cfg))
	log.Print(services.GetServerStartedLogString(time.Now(), cfg.VideoGatewayPort, "video-service"))

	log.Fatal(server.Serve(lis))
}
