package handlers

import (
	userservice "api-repository/pkg/api/user-service"
	"database/sql"
	"sync"
)

var once sync.Once

type AuthHandler struct {
	db *sql.DB
}

func NewAuthHandler(_db *sql.DB) *AuthHandler {
	var handler *AuthHandler
	once.Do(func() {
		handler = &AuthHandler{
			db: _db,
		}
	})
	return handler
}

func (a *AuthHandler) Register(req *userservice.RegisterRequest) *userservice.RegisterResponse {
	return &userservice.RegisterResponse{
		Uuid:    "1234",
		IsAdmin: false,
	}
}

func (a *AuthHandler) Login(req *userservice.LoginRequest) *userservice.LoginResponse {
	return &userservice.LoginResponse{
		Uuid: "1234",
	}
}
