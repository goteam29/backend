package handlers

import (
	"database/sql"
	"github.com/redis/go-redis/v9"
)

type TextHandler struct {
	pg    *sql.DB
	redis *redis.Client
}

func NewTextHandler(pgConn *sql.DB, redisConn *redis.Client) *TextHandler {
	return &TextHandler{
		pg:    pgConn,
		redis: redisConn,
	}
}
