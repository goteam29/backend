package main

import (
	"api-repository/internal/config"
	"api-repository/internal/services/file-service/service"
	"log"
	"os"
	"os/signal"
	"syscall"
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

	if err := fileService.Start(cfg.FileServicePort); err != nil {
		log.Fatalf("Failed to start service: %v", err)
	}
	defer fileService.Stop()

	log.Printf("Service started on port %d", cfg.FileServicePort)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
}
