package internal

import (
	"api-repository/internal/user-sevice/service/internal/handlers"
	userservice "api-repository/pkg/api/user-service"
	"database/sql"
	"sync"
)

var once sync.Once

type UserService struct {
	userservice.UnimplementedUserServer
	authHandler *handlers.AuthHandler
	pgConn      *sql.DB
}

func NewUserService(pc *sql.DB) *UserService {
	var s *UserService
	once.Do(func() {
		s = &UserService{
			pgConn:      pc,
			authHandler: handlers.NewAuthHandler(pc),
		}
	})
	return s
}

func (us *UserService) Register(request *userservice.RegisterRequest) (*userservice.RegisterResponse, error) {
	return us.authHandler.Register(request), nil
}

func (us *UserService) Login(request *userservice.LoginRequest) (*userservice.LoginResponse, error) {
	return us.authHandler.Login(request), nil
}
