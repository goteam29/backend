package main

import (
	"api-repository/internal/config"
	"api-repository/internal/services/text-service/service"
	textService "api-repository/pkg/api/text-service"
	"api-repository/pkg/db/postgres"
	"log"
	"net"
	"strconv"

	"google.golang.org/grpc"
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

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(cfg.TextServicePort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	svc := service.NewTextService(pgConn)
	textService.RegisterTextServer(server, svc)

	log.Print("Text service started at port: ", cfg.TextServicePort)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
