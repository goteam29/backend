package main

import (
	us "api-repository/pkg/api/user-service"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect to UserService: %v", err)
	}
	defer conn.Close()

	client := us.NewUserClient(conn)
	if err := us.RegisterUserHandlerClient(ctx, mux, client); err != nil {
		log.Fatalf("failed to register the user server: %v", err)
	}

	addr := ":8080"
	fmt.Println("GATEWAY STARTED | port : 8080")
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Gateway stopped: %v", err)
	}
}
