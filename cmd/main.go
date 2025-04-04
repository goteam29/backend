package main

import (
	mainConfig "api-repository/internal/config"
	userservice "api-repository/pkg/api/user-service"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	userservice.UnimplementedUserServer
}

func New() *UserService {
	return &UserService{}
}

func (u *UserService) Get(ctx context.Context, req *userservice.Request) (*userservice.Reply, error) {
	if req.Message == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request cannot be empty")
	}

	return &userservice.Reply{Message: req.Message}, nil
}

const port = "50050"

func main() {
	ctx := context.Background()
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	config, err := mainConfig.NewMainConfig()
	if err != nil {
		log.Fatalf("can't read config %v", err)
	}
	log.Println("Config:", config)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	service := New()
	userservice.RegisterUserServer(server, service)

	log.Printf("server listening at %v", lis.Addr())
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-ctx.Done()
	server.GracefulStop()
}
