package handlers

import (
	repository "api-repository/internal/services/text-service/service/internal/repository/lesson"
	textService "api-repository/pkg/api/text-service"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (th *TextHandler) CreateLesson(ctx context.Context, req *textService.CreateLessonRequest) (*textService.CreateLessonResponse, error) {
	id := uuid.New()

	err := repository.RedisAdd(id, th.redis, req)
	if err != nil {
		return nil, fmt.Errorf("createLesson: %v", err)
	}

	err = repository.PgInsert(id, th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("createLesson: %v", err)
	}

	return &textService.CreateLessonResponse{
		Response: "lesson created successfully",
	}, nil
}

func (th *TextHandler) GetLesson(ctx context.Context, req *textService.GetLessonRequest) (*textService.GetLessonResponse, error) {
	lesson, err := repository.RedisGet(th.redis, ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("getLesson: %v", err)
	}

	if lesson == nil {
		lesson, err := repository.PgSelect(th.pg, req)
		if err != nil {
			return nil, fmt.Errorf("getLesson: %v", err)
		}

		log.Print("lesson found in PostgreSQL")
		return lesson, nil
	}

	log.Print("lesson found in Redis")
	return lesson, nil
}
