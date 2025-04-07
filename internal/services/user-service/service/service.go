package service

import (
	"api-repository/internal/services/user-service/service/internal/handlers"
	userservice "api-repository/pkg/api/user-service"
	"context"
	"database/sql"
	"os"
	"sync"
)

var once sync.Once

type UserService struct {
	userservice.UnimplementedUserServer
	authHandler *handlers.AuthHandler
	pgConn      *sql.DB
}

func NewUserService(pc *sql.DB) *UserService {
	secret := os.Getenv("JWT_SECRET_TOKEN")

	var s *UserService
	once.Do(func() {
		s = &UserService{
			pgConn:      pc,
			authHandler: handlers.NewAuthHandler(pc, secret),
		}
	})
	return s
}

func (us *UserService) Register(ctx context.Context, request *userservice.RegisterRequest) (*userservice.RegisterResponse, error) {
	return us.authHandler.Register(request)
}

func (us *UserService) Login(ctx context.Context, request *userservice.LoginRequest) (*userservice.LoginResponse, error) {
	return us.authHandler.Login(request)
}
