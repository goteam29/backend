package handlers

import (
	userservice "api-repository/pkg/api/user-service"
	jwtManager "api-repository/pkg/utils"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const expirationTime = 7 * 24 * time.Hour

type AuthHandler struct {
	db     *sql.DB
	secret string
	jwt    *jwtManager.JWTManager
}

func NewAuthHandler(_db *sql.DB, _secret string) *AuthHandler {
	JWTManager := jwtManager.NewJWTManager(_secret, expirationTime)
	return &AuthHandler{
		db:     _db,
		secret: _secret,
		jwt:    JWTManager,
	}
}

func (a *AuthHandler) Register(ctx context.Context, req *userservice.RegisterRequest) (*userservice.RegisterResponse, error) {
	if req.Password != req.PasswordConfirm {
		return nil, errors.New("passwords don't match")
	}

	id, token, err := a.jwt.Generate(req)
	if err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}

	_, err = a.db.Exec("INSERT INTO users (id, username, email, token) VALUES($1, $2, $3, $4)",
		id,
		req.Username,
		req.Email,
		token,
	)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code == "23505" {
				return nil, fmt.Errorf("user with this email already exists")
			}
		} else {
			return nil, fmt.Errorf("failed to insert user into database: %w", err)
		}
	}

	err = grpc.SendHeader(ctx, metadata.Pairs(
		"Set-Cookie", "access_token="+token+"; HttpOnly; Path=/; SameSite=Lax",
	))
	if err != nil {
		return nil, err
	}

	return &userservice.RegisterResponse{
		Uuid: "user successfully registered",
	}, nil
}

func (a *AuthHandler) Login(ctx context.Context, req *userservice.LoginRequest) (*userservice.LoginResponse, error) {
	user := a.db.QueryRow("SELECT token FROM users WHERE email = $1;", req.Email)
	var token string

	if err := user.Scan(&token); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with this email doesn't exist")
		}
		return nil, fmt.Errorf("login: failed to scan user: %w", err)
	}

	claims, err := a.jwt.Verify(token)
	if err != nil {
		return nil, fmt.Errorf("failed to verify token: %w", err)
	} else {
		if claims.Password != req.Password {
			return nil, fmt.Errorf("wrong password")
		}
	}

	res := &userservice.LoginResponse{Uuid: "account logged in"}
	return res, nil
}
