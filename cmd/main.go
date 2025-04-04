package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	userService "common/protos/user_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	userService.UnimplementedUserServer
}

func New() *UserService {
	return &UserService{}
}

func (u *UserService) Get(ctx context.Context, req *userService.Request) (*userService.Reply, error) {
	if req.Message == "" {
		return nil, status.Errorf(codes.InvalidArgument, "request cannot be empty")
	}

	return &userService.Reply{Message: req.Message}, nil
}

const port = "8080"

func main() {
	ctx := context.Background()
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	service := New()
	userService.RegisterUserServer(server, service)

	log.Printf("server listening at %v", lis.Addr())
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-ctx.Done()
	server.GracefulStop()
}
