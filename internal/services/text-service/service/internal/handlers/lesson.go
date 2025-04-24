package handlers

import (
	postgresRepo "api-repository/internal/services/text-service/service/internal/repository/postgres"
	textService "api-repository/pkg/api/text-service"
	"context"
	"fmt"
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

func (th *TextHandler) AddVideoInLesson(ctx context.Context, req *textService.AddVideoInLessonRequest) (*textService.AddVideoInLessonResponse, error) {
	id, err := postgresRepo.AddVideoInLesson(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("addVideoInLesson: %v", err)
	}

	return id, nil
}

func (th *TextHandler) AddFileInLesson(ctx context.Context, req *textService.AddFileInLessonRequest) (*textService.AddFileInLessonResponse, error) {
	id, err := postgresRepo.AddFileInLesson(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("addFileInLesson: %v", err)
	}

	return id, nil
}

func (th *TextHandler) AddExerciseInLesson(ctx context.Context, req *textService.AddExerciseInLessonRequest) (*textService.AddExerciseInLessonResponse, error) {
	id, err := postgresRepo.AddExerciseInLesson(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("addExerciseInLesson: %v", err)
	}

	return id, nil
}

func (th *TextHandler) AddCommentInLesson(ctx context.Context, req *textService.AddCommentInLessonRequest) (*textService.AddCommentInLessonResponse, error) {
	id, err := postgresRepo.AddCommentInLesson(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("addCommentInLesson: %v", err)
	}

	return id, nil
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

func (th *TextHandler) DeleteLesson(ctx context.Context, req *textService.DeleteLessonRequest) (*textService.DeleteLessonResponse, error) {
	id, err := postgresRepo.DeleteLesson(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("deleteLesson: %v", err)
	}

	return id, nil
}
