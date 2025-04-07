package handlers

import (
	userservice "api-repository/pkg/api/user-service"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
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

func (a *AuthHandler) Register(ctx context.Context, req *userservice.RegisterRequest) (*userservice.RegisterResponse, error) {
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

	userRow := a.db.QueryRow("SELECT id FROM users WHERE email = $1", req.Email)

	err = userRow.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			_, err = a.db.Exec("INSERT INTO users VALUES($1, $2, $3, $4)",
				id,
				req.Username,
				req.Email,
				token,
			)
			if err != nil {
				return nil, fmt.Errorf("failed to insert user: %w", err)
			}
		} else {
			return nil, fmt.Errorf("failed to query user: %w", err)
		}
	} else {
		return nil, errors.New("user already exists")
	}

	err = grpc.SendHeader(ctx, metadata.Pairs(
		"Set-Cookie", "access_token="+token+"; HttpOnly; Path=/; SameSite=Lax",
	))
	if err != nil {
		return nil, err
	}

	return &userservice.RegisterResponse{
		Uuid:    id.String(),
		IsAdmin: false,
	}, nil
}

func (a *AuthHandler) Login(ctx context.Context, req *userservice.LoginRequest) (*userservice.LoginResponse, error) {
	return &userservice.LoginResponse{
		Uuid: "1234",
	}, nil
}
