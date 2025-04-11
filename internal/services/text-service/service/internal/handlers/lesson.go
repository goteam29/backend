package handlers

import (
	postgresRepo "api-repository/internal/services/text-service/service/internal/repository/postgres"
	redisRepo "api-repository/internal/services/text-service/service/internal/repository/redis"
	textService "api-repository/pkg/api/text-service"
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
)

func (th *TextHandler) CreateLesson(ctx context.Context, req *textService.CreateLessonRequest) (*textService.CreateLessonResponse, error) {
	id := uuid.New()

	err := redisRepo.AddLesson(id, th.redis, req)
	if err != nil {
		return nil, fmt.Errorf("createLesson: %v", err)
	}

	err = postgresRepo.InsertLesson(id, th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("createLesson: %v", err)
	}

	return &textService.CreateLessonResponse{
		Response: "lesson created successfully",
	}, nil
}

func (th *TextHandler) GetLesson(ctx context.Context, req *textService.GetLessonRequest) (*textService.GetLessonResponse, error) {
	lesson, err := redisRepo.GetLesson(th.redis, ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("getLesson: %v", err)
	}

	if lesson == nil {
		lesson, err := postgresRepo.SelectLesson(th.pg, req)
		if err != nil {
			return nil, fmt.Errorf("getLesson: %v", err)
		}

		log.Print("lesson found in PostgreSQL")
		return lesson, nil
	}

	log.Print("lesson found in Redis")
	return lesson, nil
}
