package handlers

import (
	postgresRepo "api-repository/internal/services/text-service/service/internal/repository/postgres"
	textService "api-repository/pkg/api/text-service"
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (th *TextHandler) CreateLesson(ctx context.Context, req *textService.CreateLessonRequest) (*textService.CreateLessonResponse, error) {
	id, err := postgresRepo.InsertLesson(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("createLesson: %v", err)
	}

	return id, nil
}

func (th *TextHandler) GetLesson(ctx context.Context, req *textService.GetLessonRequest) (*textService.GetLessonResponse, error) {
	lesson, err := postgresRepo.SelectLesson(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("getLesson: %v", err)
	}

	return lesson, nil
}

func (th *TextHandler) GetLessons(ctx context.Context, req *textService.GetLessonsRequest) (*textService.GetLessonsResponse, error) {
	lessons, err := postgresRepo.SelectLessons(th.pg)
	if err != nil {
		return nil, fmt.Errorf("getLessons: %v", err)
	}

	return lessons, nil
}

func (th *TextHandler) IncreaseRating(ctx context.Context, req *textService.IncreaseRatingRequest) (*textService.IncreaseRatingResponse, error) {
	id, err := postgresRepo.IncreaseRating(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("increaseRating: %v", err)
	}

	return id, nil
}

func (th *TextHandler) DecreaseRating(ctx context.Context, req *textService.DecreaseRatingRequest) (*textService.DecreaseRatingResponse, error) {
	id, err := postgresRepo.DecreaseRating(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("decreaseRating: %v", err)
	}

	return id, nil
}

func (th *TextHandler) DeleteLesson(ctx context.Context, req *textService.DeleteLessonRequest) (*emptypb.Empty, error) {
	id, err := postgresRepo.DeleteLesson(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("deleteLesson: %v", err)
	}

	return id, nil
}
