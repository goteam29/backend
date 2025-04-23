package handlers

import (
	userservice "api-repository/pkg/api/user-service"
	jwtManager "api-repository/pkg/utils"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
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

func (ah *AuthHandler) Register(ctx context.Context, req *userservice.RegisterRequest) (*userservice.RegisterResponse, error) {
	if req.Password != req.PasswordConfirm {
		return nil, status.Error(codes.InvalidArgument, "passwords do not match")
	}

	log.Printf("%v", req)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't hash password %v", err)
	}

	id := uuid.New().String()

	query := `INSERT into users (id, username, email, password_hash) values ($1, $2, $3, $4)`
	_, err = ah.db.Exec(query, id, req.Username, req.Email, string(passwordHash))
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" {
				return nil, fmt.Errorf("user with this email already exists")
			}
		}
		return nil, status.Errorf(codes.Internal, "can't insert user into bd | err: %v", err)
	}

	token, err := ah.jwt.Generate(id, req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "can't generate jwt token | err %v", err)
	}

	return &userservice.RegisterResponse{
		Token: token,
		Uuid:  id,
	}, nil
}

func (a *AuthHandler) Login(ctx context.Context, req *userservice.LoginRequest) (*userservice.LoginResponse, error) {
	row := a.db.QueryRow(`SELECT id, password_hash FROM users WHERE email = $1`, req.Email)

	var id string
	var passwordHash string
	if err := row.Scan(&id, &passwordHash); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to query user: %v", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(req.Password)); err != nil {
		return nil, status.Error(codes.Unauthenticated, "incorrect password")
	}

	token, err := a.jwt.Generate(id, req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate token: %v", err)
	}

	return &userservice.LoginResponse{
		Uuid:  id,
		Token: token,
	}, nil
}
