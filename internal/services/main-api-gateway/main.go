package main

import (
	fs "api-repository/pkg/api/file-service"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fileServiceAddr := "50052"
	conn, err := grpc.NewClient(fileServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("err: | %v", err)
	}
	defer conn.Close()

	mux := runtime.NewServeMux()
	if err := fs.RegisterFileHandlerClient(context.Background(), mux, conn); err != nil {
		log.Fatalf("failed to register the file server: %v", err)
	}

	addr := "0.0.0.0:8080"
	fmt.Println("API gateway server is running on " + addr)
	if err = http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("gateway server closed abruptly: ", err)
	}

}
