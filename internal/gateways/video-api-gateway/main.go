package main

import (
	"api-repository/internal/config"
	videoService "api-repository/pkg/api/video-service"
	"api-repository/pkg/utils"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

func main() {
	utils.CreateNewSugaredLogger()

	cfg, err := config.NewMainConfig()
	if err != nil {
		utils.GetSugaredLogger().Fatalf("can't get config | err: %v", err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()

	createVideoServiceConnection(ctx, mux, cfg)

	utils.GetSugaredLogger().Logf(0, "VIDEO-GATEWAY STARTED | port %d", cfg.VideoGatewayPort)
	if err = http.ListenAndServe(fmt.Sprintf(":%d", cfg.VideoGatewayPort), mux); err != nil {
		utils.GetSugaredLogger().Fatalf("Gateway stopped %v", err)
	}

}

func createVideoServiceConnection(ctx context.Context, mux *runtime.ServeMux, cfg *config.MainConfig) {
	conn, err := grpc.NewClient(fmt.Sprintf("localhost:%d", cfg.VideServicePort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		utils.GetSugaredLogger().Fatalf("can't connect to video-service | err: %v", err)
	}

	client := videoService.NewVideoClient(conn)
	if err = videoService.RegisterVideoHandlerClient(ctx, mux, client); err != nil {
		utils.GetSugaredLogger().Fatalf("failed to register video-service client | err: %v", err)
	}
}
