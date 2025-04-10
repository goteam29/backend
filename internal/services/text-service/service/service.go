package service

import (
	"api-repository/internal/services/text-service/service/internal/handlers"
	textService "api-repository/pkg/api/text-service"
	"context"
	"database/sql"
	"sync"

	"github.com/redis/go-redis/v9"
)

var once sync.Once

type TextService struct {
	textService.UnimplementedTextServer
	textHandler *handlers.TextHandler
	pgConn      *sql.DB
}

func NewTextService(pg *sql.DB, redis *redis.Client) *TextService {
	var s *TextService
	once.Do(func() {
		s = &TextService{
			pgConn:      pg,
			textHandler: handlers.NewTextHandler(pg, redis),
		}
	})

	return s
}

func (ts *TextService) CreateLesson(ctx context.Context, request *textService.CreateLessonRequest) (*textService.CreateLessonResponse, error) {
	return ts.textHandler.CreateLesson(ctx, request)
}

func (ts *TextService) GetLesson(ctx context.Context, request *textService.GetLessonRequest) (*textService.GetLessonResponse, error) {
	return ts.textHandler.GetLesson(ctx, request)
}
