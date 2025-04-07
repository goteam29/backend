package main

import (
	"api-repository/internal/config"
	"api-repository/internal/file-service/service"
	"log"
)

func main() {
	cfg, err := config.NewMainConfig()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	fileService, err := service.New(cfg)
	if err != nil {
		log.Fatalf("Service initialization failed: %v", err)
	}

	log.Printf("Starting gRPC server on port %d", cfg.UserServicePort)
	if err := fileService.Start(cfg.UserServicePort); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
