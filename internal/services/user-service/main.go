package main

import (
	"api-repository/internal/config"
	"api-repository/internal/services"
	"api-repository/internal/services/user-service/service"
	userservice "api-repository/pkg/api/user-service"
	"api-repository/pkg/db/postgres"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	cfg, err := config.NewMainConfig()
	if err != nil {
		log.Fatalf("can't get env files | err: %v", err)
	}

	pgConn, err := postgres.NewPostgres(cfg.POSTGRES)
	if err != nil {
		log.Fatalf("can't connect to postgres | err: %v", err)
	}
	svc := service.NewUserService(pgConn)

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(cfg.UserServicePort))
	if err != nil {
		log.Fatalf("can't start servier | err: %v", err)
	}

	server := grpc.NewServer()
	userservice.RegisterUserServer(server, svc)

	log.Print(services.GetServerStartedLogString(cfg, time.Now(), cfg.UserServicePort, "user-service"))
	log.Fatal(server.Serve(lis))
}
