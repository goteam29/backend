package utils

import (
	userservice "api-repository/pkg/api/user-service"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

type UserClaims struct {
	jwt.RegisteredClaims
	UserID   uuid.UUID `json:"userid"`
	Email    string    `json:"name"`
	Password string    `json:"password"`
}

func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{secretKey, tokenDuration}
}

func (manager *JWTManager) Generate(req *userservice.RegisterRequest) (uuid.UUID, string, error) {
	id := uuid.New()

	claims := UserClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(manager.tokenDuration)),
		},
		UserID:   id,
		Email:    req.Email,
		Password: req.Password,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(manager.secretKey))
	if err != nil {
		return uuid.Nil, "", fmt.Errorf("failed to sign token: %w", err)
	}

	return id, tokenStr, nil
}

func (manager *JWTManager) Verify(jwtToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		jwtToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("unexpected token signing method")
			}

			return []byte(manager.secretKey), nil
		},
	)
	if err != nil {
		return nil, fmt.Errorf("can't parse token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
