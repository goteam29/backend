package main

import (
	"api-repository/internal/config"
	"api-repository/internal/services/text-service/service"
	textService "api-repository/pkg/api/text-service"
	"api-repository/pkg/db/postgres"
	"api-repository/pkg/db/redis"
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"

	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	cfg, err := config.NewMainConfig()
	if err != nil {
		log.Fatalf("can't get env files | err: %v", err)
	}

	pgConn, err := postgres.NewPostgres(cfg.POSTGRES)
	if err != nil {
		log.Fatalf("can't connect to postgres | err: %v", err)
	}

	redisConn := redis.NewRedisConn(cfg.REDIS)
	_, err = redisConn.Ping(ctx).Result()
	if err != nil {
		log.Printf("can't connect to redis | err: %v", err)
	}

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(cfg.TextServicePort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	svc := service.NewTextService(pgConn, redisConn)
	textService.RegisterTextServer(server, svc)

	log.Print("text service started at port: ", cfg.TextServicePort)
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-ctx.Done()
	server.GracefulStop()
	log.Print("text service stopped")
}
