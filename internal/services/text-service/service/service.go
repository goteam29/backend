package service

import (
	"api-repository/internal/services/text-service/service/internal/handlers"
	textService "api-repository/pkg/api/text-service"
	"context"
	"database/sql"

	"github.com/redis/go-redis/v9"
)

type TextService struct {
	textService.UnimplementedTextServer
	pgConn      *sql.DB
	textHandler *handlers.TextHandler
}

func NewTextService(pg *sql.DB, rds *redis.Client) *TextService {
	return &TextService{
		pgConn:      pg,
		textHandler: handlers.NewTextHandler(pg, rds),
	}
}

func (ts *TextService) CreateClass(ctx context.Context, request *textService.CreateClassRequest) (*textService.CreateClassResponse, error) {
	return ts.textHandler.CreateClass(ctx, request)
}

func (ts *TextService) GetClass(ctx context.Context, request *textService.GetClassRequest) (*textService.GetClassResponse, error) {
	return ts.textHandler.GetClass(ctx, request)
}

func (ts *TextService) GetClasses(ctx context.Context, request *textService.GetClassesRequest) (*textService.GetClassesResponse, error) {
	return ts.textHandler.GetClasses(ctx)
}

func (ts *TextService) CreateSubject(ctx context.Context, request *textService.CreateSubjectRequest) (*textService.CreateSubjectResponse, error) {
	return ts.textHandler.CreateSubject(ctx, request)
}

func (ts *TextService) CreateLesson(ctx context.Context, request *textService.CreateLessonRequest) (*textService.CreateLessonResponse, error) {
	return ts.textHandler.CreateLesson(ctx, request)
}

func (ts *TextService) GetLesson(ctx context.Context, request *textService.GetLessonRequest) (*textService.GetLessonResponse, error) {
	return ts.textHandler.GetLesson(ctx, request)
}
