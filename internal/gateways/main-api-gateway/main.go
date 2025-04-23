package main

import (
	"api-repository/internal/config"
	fileservice "api-repository/pkg/api/file-service"
	textService "api-repository/pkg/api/text-service"
	us "api-repository/pkg/api/user-service"
	"api-repository/pkg/utils"
	"context"
	"fmt"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func main() {
	utils.CreateNewSugaredLogger()

	cfg, err := config.NewMainConfig()
	if err != nil {
		utils.GetSugaredLogger().Logf(zapcore.InfoLevel, "can't get config | err: %v", err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	createUserServiceConnection(ctx, mux, cfg)
	createTextServiceConnection(ctx, mux, cfg)
	createFileServiceConnection(ctx, mux, cfg)

	utils.GetSugaredLogger().Logf(0, "GATEWAY STARTED | port: %d", cfg.GatewayPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.GatewayPort), mux); err != nil {
		log.Fatalf("Gateway stopped: %v", err)
	}
}

func createUserServiceConnection(ctx context.Context, mux *runtime.ServeMux, cfg *config.MainConfig) {
	conn, err := grpc.NewClient(fmt.Sprintf("user-service:%d", cfg.UserServicePort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		utils.GetSugaredLogger().Fatalf("can't connect to user-service | err: %v", err)
	}

	client := us.NewUserClient(conn)
	if err = us.RegisterUserHandlerClient(ctx, mux, client); err != nil {
		utils.GetSugaredLogger().Fatalf("failed to register user-service client | err: %v", err)
	}
}

func createTextServiceConnection(ctx context.Context, mux *runtime.ServeMux, cfg *config.MainConfig) {
	conn, err := grpc.NewClient(fmt.Sprintf("text-service:%d", cfg.TextServicePort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		utils.GetSugaredLogger().Fatalf("can't connect to text-service | err: %v", err)
	}

	client := textService.NewTextClient(conn)
	if err = textService.RegisterTextHandlerClient(ctx, mux, client); err != nil {
		utils.GetSugaredLogger().Fatalf("failed to register text handler client | err: %v", err)
	}
}

func createFileServiceConnection(ctx context.Context, mux *runtime.ServeMux, cfg *config.MainConfig) {
	conn, err := grpc.NewClient(fmt.Sprintf("file-service:%d", cfg.FileServicePort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		utils.GetSugaredLogger().Fatalf("can' connect to file-service | err: %v", err)
	}

	client := fileservice.NewFileClient(conn)
	if err = fileservice.RegisterFileHandlerClient(ctx, mux, client); err != nil {
		utils.GetSugaredLogger().Fatalf("failed to register file handler client | err: %v", err)
	}
}
