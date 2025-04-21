package handlers

import (
	postgresRepo "api-repository/internal/services/text-service/service/internal/repository/postgres"
	textService "api-repository/pkg/api/text-service"
	"context"
	"fmt"
)

func (th *TextHandler) CreateLesson(ctx context.Context, req *textService.CreateLessonRequest) (*textService.CreateLessonResponse, error) {
	// err := redisRepo.AddLesson(id, th.redis, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("createLesson: %v", err)
	// }

	id, err := postgresRepo.InsertLesson(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("createLesson: %v", err)
	}

	return id, nil
}

func (th *TextHandler) GetLesson(ctx context.Context, req *textService.GetLessonRequest) (*textService.GetLessonResponse, error) {
	// lesson, err := redisRepo.GetLesson(th.redis, ctx, req.Id)
	// if err != nil {
	// 	return nil, fmt.Errorf("getLesson: %v", err)
	// }

	// if lesson == nil {
	lesson, err := postgresRepo.SelectLesson(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("getLesson: %v", err)
	}

	return lesson, nil
	// }

	// log.Print("lesson found in Redis")
	// return lesson, nil
}

func (th *TextHandler) GetLessons(ctx context.Context, req *textService.GetLessonsRequest) (*textService.GetLessonsResponse, error) {
	// lessons, err := redisRepo.GetLessons(th.redis, ctx, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("getLessons: %v", err)
	// }

	lessons, err := postgresRepo.SelectLessons(th.pg)
	if err != nil {
		return nil, fmt.Errorf("getLessons: %v", err)
	}

	return lessons, nil
}

func (th *TextHandler) AddVideoInLesson(ctx context.Context, req *textService.AddVideoInLessonRequest) (*textService.AddVideoInLessonResponse, error) {
	// err := redisRepo.AddVideoInLesson(th.redis, ctx, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("addVideoInLesson: %v", err)
	// }

	id, err := postgresRepo.AddVideoInLesson(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("addVideoInLesson: %v", err)
	}

	return id, nil
}

func (th *TextHandler) RemoveVideoFromLesson(ctx context.Context, req *textService.RemoveVideoFromLessonRequest) (*textService.RemoveVideoFromLessonResponse, error) {
	// err := redisRepo.RemoveVideoFromLesson(th.redis, ctx, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("removeVideoFromLesson: %v", err)
	// }

	id, err := postgresRepo.RemoveVideoFromLesson(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("removeVideoFromLesson: %v", err)
	}

	return id, nil
}

func (th *TextHandler) AddFileInLesson(ctx context.Context, req *textService.AddFileInLessonRequest) (*textService.AddFileInLessonResponse, error) {
	// err := redisRepo.AddFileInLesson(th.redis, ctx, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("addFileInLesson: %v", err)
	// }

	id, err := postgresRepo.AddFileInLesson(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("addFileInLesson: %v", err)
	}

	return id, nil
}

func (th *TextHandler) RemoveFileFromLesson(ctx context.Context, req *textService.RemoveFileFromLessonRequest) (*textService.RemoveFileFromLessonResponse, error) {
	// err := redisRepo.RemoveFileFromLesson(th.redis, ctx, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("removeFileFromLesson: %v", err)
	// }

	id, err := postgresRepo.RemoveFileFromLesson(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("removeFileFromLesson: %v", err)
	}

	return id, nil
}

func (th *TextHandler) AddExerciseInLesson(ctx context.Context, req *textService.AddExerciseInLessonRequest) (*textService.AddExerciseInLessonResponse, error) {
	// err := redisRepo.AddExerciseInLesson(th.redis, ctx, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("addExerciseInLesson: %v", err)
	// }

	id, err := postgresRepo.AddExerciseInLesson(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("addExerciseInLesson: %v", err)
	}

	return id, nil
}

func (th *TextHandler) RemoveExerciseFromLesson(ctx context.Context, req *textService.RemoveExerciseFromLessonRequest) (*textService.RemoveExerciseFromLessonResponse, error) {
	// err := redisRepo.RemoveExerciseFromLesson(th.redis, ctx, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("removeExerciseFromLesson: %v", err)
	// }

	id, err := postgresRepo.RemoveExerciseFromLesson(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("removeExerciseFromLesson: %v", err)
	}

	return id, nil
}

func (th *TextHandler) AddCommentInLesson(ctx context.Context, req *textService.AddCommentInLessonRequest) (*textService.AddCommentInLessonResponse, error) {
	// err := redisRepo.AddCommentInLesson(th.redis, ctx, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("addCommentInLesson: %v", err)
	// }

	id, err := postgresRepo.AddCommentInLesson(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("addCommentInLesson: %v", err)
	}

	return id, nil
}

func (th *TextHandler) RemoveCommentFromLesson(ctx context.Context, req *textService.RemoveCommentFromLessonRequest) (*textService.RemoveCommentFromLessonResponse, error) {
	// err := redisRepo.RemoveCommentFromLesson(th.redis, ctx, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("removeCommentFromLesson: %v", err)
	// }

	id, err := postgresRepo.RemoveCommentFromLesson(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("removeCommentFromLesson: %v", err)
	}

	return id, nil
}

func (th *TextHandler) IncreaseRating(ctx context.Context, req *textService.IncreaseRatingRequest) (*textService.IncreaseRatingResponse, error) {
	// err := redisRepo.IncreaseRating(th.redis, ctx, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("increaseRating: %v", err)
	// }

	id, err := postgresRepo.IncreaseRating(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("increaseRating: %v", err)
	}

	return id, nil
}

func (th *TextHandler) DecreaseRating(ctx context.Context, req *textService.DecreaseRatingRequest) (*textService.DecreaseRatingResponse, error) {
	// err := redisRepo.DecreaseRating(th.redis, ctx, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("decreaseRating: %v", err)
	// }

	id, err := postgresRepo.DecreaseRating(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("decreaseRating: %v", err)
	}

	return id, nil
}

func (th *TextHandler) DeleteLesson(ctx context.Context, req *textService.DeleteLessonRequest) (*textService.DeleteLessonResponse, error) {
	// err := redisRepo.DeleteLesson(th.redis, ctx, req.Id)
	// if err != nil {
	// 	return nil, fmt.Errorf("deleteLesson: %v", err)
	// }

	id, err := postgresRepo.DeleteLesson(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("deleteLesson: %v", err)
	}

	return id, nil
}
