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

	id := uuid.New()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":   req.Username,
		"userid": id,
		"exp":    time.Now().Add(time.Hour * 24 * 7).Unix(),
		"iat":    time.Now().Unix(),
	})

	token, err := claims.SignedString([]byte(a.secret))
	if err != nil {
		return nil, err
	}

	_, err = a.db.Exec("INSERT INTO users VALUES($1, $2, $3, $4) ON CONFLICT (email) DO NOTHING ",
		id,
		req.Username,
		req.Email,
		token,
	)
	if err != nil {
		return nil, err
	}

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
