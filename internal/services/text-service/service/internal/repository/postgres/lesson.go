package repository

import (
	textService "api-repository/pkg/api/text-service"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func InsertLesson(pg *sql.DB, req *textService.CreateLessonRequest) (*textService.CreateLessonResponse, error) {
	id := uuid.New()

	_, err := pg.Exec("INSERT INTO lessons (id, section_id, name, description) VALUES ($1, $2, $3, $4)",
		id,
		req.SectionId,
		req.Name,
		req.Description,
	)
	if err != nil {
		return nil, fmt.Errorf("pgInsertLesson: failed to insert lesson into database: %v", err)
	}

	return &textService.CreateLessonResponse{
		Id: id.String(),
	}, nil
}

func SelectLesson(pg *sql.DB, req *textService.GetLessonRequest) (*textService.GetLessonResponse, error) {
	lessonResponse := &textService.GetLessonResponse{
		Lesson: &textService.Lesson{
			VideoIds:    make([]string, 0),
			FileIds:     make([]string, 0),
			ExerciseIds: make([]string, 0),
			CommentIds:  make([]string, 0),
		},
	}

	lesson := pg.QueryRow("SELECT id, section_id, name, description FROM lessons WHERE id = $1", req.Id)
	err := lesson.Scan(
		&lessonResponse.Lesson.Id, &lessonResponse.Lesson.SectionId,
		&lessonResponse.Lesson.Name, &lessonResponse.Lesson.Description)
	if err != nil {
		return nil, fmt.Errorf("pgSelectLesson: failed to scan lesson: %v", err)
	}

	return &textService.GetLessonResponse{
		Lesson: lessonResponse.Lesson,
	}, nil
}

func SelectLessons(pg *sql.DB) (*textService.GetLessonsResponse, error) {
	lessons, err := pg.Query("SELECT id, section_id, name, description, video_ids, file_ids, exercise_ids, comment_ids, rating FROM lessons")
	if err != nil {
		return nil, fmt.Errorf("pgSelectLessons: failed to select lessons from database: %v", err)
	}
	defer lessons.Close()

	lessonsResponse := make([]*textService.Lesson, 0)

	for lessons.Next() {
		lesson := &textService.Lesson{}
		var videoIds, fileIds, exerciseIds, commentIds pq.StringArray

		err := lessons.Scan(&lesson.Id, &lesson.SectionId, &lesson.Name, &lesson.Description,
			&videoIds, &fileIds, &exerciseIds, &commentIds, &lesson.Rating)
		if err != nil {
			return nil, fmt.Errorf("pgSelectLessons: failed to scan lesson: %v", err)
		}

		lesson.VideoIds = videoIds
		lesson.FileIds = fileIds
		lesson.ExerciseIds = exerciseIds
		lesson.CommentIds = commentIds

		lessonsResponse = append(lessonsResponse, lesson)
	}

	if err := lessons.Err(); err != nil {
		return nil, fmt.Errorf("pgSelectLessons: error during rows iteration: %v", err)
	}

	return &textService.GetLessonsResponse{
		Lessons: lessonsResponse,
	}, nil
}

func AddVideoInLesson(pg *sql.DB, req *textService.AddVideoInLessonRequest) (*textService.AddVideoInLessonResponse, error) {
	_, err := pg.Exec("UPDATE lessons SET video_ids = array_append(video_ids, $1) WHERE id = $2", req.VideoId, req.Id)
	if err != nil {
		return nil, fmt.Errorf("pgAddVideoInLesson: failed to add video in lesson: %v", err)
	}

	return &textService.AddVideoInLessonResponse{
		VideoId: req.VideoId,
	}, nil
}

func RemoveVideoFromLesson(pg *sql.DB, req *textService.RemoveVideoFromLessonRequest) (*textService.RemoveVideoFromLessonResponse, error) {
	_, err := pg.Exec("UPDATE lessons SET video_ids = array_remove(video_ids, $1) WHERE id = $2", req.VideoId, req.Id)
	if err != nil {
		return nil, fmt.Errorf("pgRemoveVideoFromLesson: failed to remove video from lesson: %v", err)
	}

	return &textService.RemoveVideoFromLessonResponse{
		VideoId: req.VideoId,
	}, nil
}

func AddFileInLesson(pg *sql.DB, req *textService.AddFileInLessonRequest) (*textService.AddFileInLessonResponse, error) {
	_, err := pg.Exec("UPDATE lessons SET file_ids = array_append(file_ids, $1) WHERE id = $2", req.FileId, req.Id)
	if err != nil {
		return nil, fmt.Errorf("pgAddFileInLesson: failed to add file in lesson: %v", err)
	}

	return &textService.AddFileInLessonResponse{
		FileId: req.FileId,
	}, nil
}

func RemoveFileFromLesson(pg *sql.DB, req *textService.RemoveFileFromLessonRequest) (*textService.RemoveFileFromLessonResponse, error) {
	_, err := pg.Exec("UPDATE lessons SET file_ids = array_remove(file_ids, $1) WHERE id = $2", req.FileId, req.Id)
	if err != nil {
		return nil, fmt.Errorf("pgRemoveFileFromLesson: failed to remove file from lesson: %v", err)
	}

	return &textService.RemoveFileFromLessonResponse{
		FileId: req.FileId,
	}, nil
}

func AddExerciseInLesson(pg *sql.DB, req *textService.AddExerciseInLessonRequest) (*textService.AddExerciseInLessonResponse, error) {
	_, err := pg.Exec("UPDATE lessons SET exercise_ids = array_append(exercise_ids, $1) WHERE id = $2", req.ExerciseId, req.Id)
	if err != nil {
		return nil, fmt.Errorf("pgAddExerciseInLesson: failed to add exercise in lesson: %v", err)
	}

	return &textService.AddExerciseInLessonResponse{
		ExerciseId: req.ExerciseId,
	}, nil
}

func RemoveExerciseFromLesson(pg *sql.DB, req *textService.RemoveExerciseFromLessonRequest) (*textService.RemoveExerciseFromLessonResponse, error) {
	_, err := pg.Exec("UPDATE lessons SET exercise_ids = array_remove(exercise_ids, $1) WHERE id = $2", req.ExerciseId, req.Id)
	if err != nil {
		return nil, fmt.Errorf("pgRemoveExerciseFromLesson: failed to remove exercise from lesson: %v", err)
	}

	return &textService.RemoveExerciseFromLessonResponse{
		ExerciseId: req.ExerciseId,
	}, nil
}

func AddCommentInLesson(pg *sql.DB, req *textService.AddCommentInLessonRequest) (*textService.AddCommentInLessonResponse, error) {
	_, err := pg.Exec("UPDATE lessons SET comment_ids = array_append(comment_ids, $1) WHERE id = $2", req.CommentId, req.Id)
	if err != nil {
		return nil, fmt.Errorf("pgAddCommentInLesson: failed to add comment in lesson: %v", err)
	}

	return &textService.AddCommentInLessonResponse{
		CommentId: req.CommentId,
	}, nil
}

func RemoveCommentFromLesson(pg *sql.DB, req *textService.RemoveCommentFromLessonRequest) (*textService.RemoveCommentFromLessonResponse, error) {
	_, err := pg.Exec("UPDATE lessons SET comment_ids = array_remove(comment_ids, $1) WHERE id = $2", req.CommentId, req.Id)
	if err != nil {
		return nil, fmt.Errorf("pgRemoveCommentFromLesson: failed to remove comment from lesson: %v", err)
	}

	return &textService.RemoveCommentFromLessonResponse{
		CommentId: req.CommentId,
	}, nil
}

func IncreaseRating(pg *sql.DB, req *textService.IncreaseRatingRequest) (*textService.IncreaseRatingResponse, error) {
	_, err := pg.Exec("UPDATE lessons SET rating = rating + 1 WHERE id = $1", req.Id)
	if err != nil {
		return nil, fmt.Errorf("pgIncreaseRating: failed to increase lesson rating: %v", err)
	}

	return &textService.IncreaseRatingResponse{
		Id: req.Id,
	}, nil
}

func DecreaseRating(pg *sql.DB, req *textService.DecreaseRatingRequest) (*textService.DecreaseRatingResponse, error) {
	_, err := pg.Exec("UPDATE lessons SET rating = rating - 1 WHERE id = $1", req.Id)
	if err != nil {
		return nil, fmt.Errorf("pgDecreaseRating: failed to decrease lesson rating: %v", err)
	}

	return &textService.DecreaseRatingResponse{
		Id: req.Id,
	}, nil
}

func DeleteLesson(pg *sql.DB, req *textService.DeleteLessonRequest) (*textService.DeleteLessonResponse, error) {
	_, err := pg.Exec("DELETE FROM lessons WHERE id = $1", req.Id)
	if err != nil {
		return nil, fmt.Errorf("pgDeleteLesson: failed to delete lesson from database: %v", err)
	}

	return &textService.DeleteLessonResponse{
		Id: req.Id,
	}, nil
}
