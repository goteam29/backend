package handlers

import (
	userservice "api-repository/pkg/api/user-service"
	"database/sql"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

type AuthHandler struct {
	db     *sql.DB
	secret string
}

func NewAuthHandler(_db *sql.DB, _secret string) *AuthHandler {
	return &AuthHandler{
		db:     _db,
		secret: _secret,
	}
}

func (a *AuthHandler) Register(req *userservice.RegisterRequest) (*userservice.RegisterResponse, error) {
	if req.Password != req.PasswordConfirm {
		return nil, errors.New("passwords don't match")
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": req.Username,
		"exp":  time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iat":  time.Now().Unix(),
	})

	token, err := claims.SignedString([]byte(a.secret))
	if err != nil {
		return nil, err
	}
	id := uuid.New()

	a.db.Exec("INSERT ($1, $2, $3) IF NOT EXISTS ")

	return &userservice.RegisterResponse{
		Uuid:    id.String(),
		IsAdmin: false,
	}, nil
}

func (a *AuthHandler) Login(req *userservice.LoginRequest) (*userservice.LoginResponse, error) {
	return &userservice.LoginResponse{
		Uuid: "1234",
	}, nil
}
