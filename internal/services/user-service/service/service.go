package service

import (
	"api-repository/internal/config"
	"api-repository/internal/services/user-service/service/internal/handlers"
	userservice "api-repository/pkg/api/user-service"
	"context"
	"database/sql"
	"sync"
)

var once sync.Once

type UserService struct {
	userservice.UnimplementedUserServer
	authHandler *handlers.AuthHandler
	pgConn      *sql.DB
}

func NewUserService(pc *sql.DB, cfg *config.MainConfig) *UserService {
	var s *UserService
	once.Do(func() {
		s = &UserService{
			pgConn:      pc,
			authHandler: handlers.NewAuthHandler(pc, cfg.SecretToken),
		}
	})

	return s
}

func (us *UserService) Register(ctx context.Context, request *userservice.RegisterRequest) (*userservice.RegisterResponse, error) {
	return us.authHandler.Register(ctx, request)
}

func (us *UserService) Login(ctx context.Context, request *userservice.LoginRequest) (*userservice.LoginResponse, error) {
	return us.authHandler.Login(ctx, request)
}
