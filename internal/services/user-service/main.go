package main

import (
	"api-repository/internal/adapters/interceptors"
	"api-repository/internal/config"
	"api-repository/internal/services"
	"api-repository/internal/services/user-service/service"
	userservice "api-repository/pkg/api/user-service"
	"api-repository/pkg/db/postgres"
	"api-repository/pkg/utils"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

func main() {
	utils.CreateNewSugaredLogger()

	cfg, err := config.NewMainConfig()
	if err != nil {
		log.Fatalf("can't get env files | err: %v", err)
	}

	pgConn, err := postgres.NewPostgres(cfg.POSTGRES)
	if err != nil {
		log.Fatalf("can't connect to postgres | err: %v", err)
	}
	svc := service.NewUserService(pgConn, cfg)

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(cfg.UserServicePort))
	if err != nil {
		log.Fatalf("can't start servier | err: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptors.LoggingInterceptor(utils.GetSugaredLogger())),
	}
	server := grpc.NewServer(opts...)
	userservice.RegisterUserServer(server, svc)

	log.Printf("Configuration:\n%s", services.GetBeautifulConfigurationString(cfg))
	log.Print(services.GetServerStartedLogString(time.Now(), cfg.UserServicePort, "user-service"))

	log.Fatal(server.Serve(lis))
}
