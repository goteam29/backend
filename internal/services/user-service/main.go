package main

import (
	"api-repository/internal/config"
	"api-repository/internal/services"
	"api-repository/internal/services/user-service/service"
	user_service "api-repository/pkg/api/user-service"
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

	server := grpc.NewServer()
	user_service.RegisterUserServer(server, svc)

	log.Print(services.GetServerStartedLogString(time.Now(), cfg.UserServicePort, "user-service"))
	log.Printf("Configuration:\n%s", services.GetBeautifulConfigurationString(cfg))

	log.Fatal(server.Serve(lis))
}
