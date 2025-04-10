package service

import (
	"api-repository/internal/services/text-service/service/internal/handlers"
	textService "api-repository/pkg/api/text-service"
	"context"
	"database/sql"
	"sync"
)

var once sync.Once

type TextService struct {
	textService.UnimplementedTextServer
	textHandler *handlers.TextHandler
	pgConn      *sql.DB
}

func NewTextService(pc *sql.DB) *TextService {
	var s *TextService
	once.Do(func() {
		s = &TextService{
			pgConn:      pc,
			textHandler: handlers.NewTextHandler(pc),
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
