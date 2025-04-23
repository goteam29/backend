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

// Class
// Create

func (ts *TextService) CreateClass(ctx context.Context, request *textService.CreateClassRequest) (*textService.CreateClassResponse, error) {
	return ts.textHandler.CreateClass(ctx, request)
}

// Read

func (ts *TextService) GetClass(ctx context.Context, request *textService.GetClassRequest) (*textService.GetClassResponse, error) {
	return ts.textHandler.GetClass(ctx, request)
}

func (ts *TextService) GetClasses(ctx context.Context, request *textService.GetClassesRequest) (*textService.GetClassesResponse, error) {
	return ts.textHandler.GetClasses(ctx)
}

// Update

func (ts *TextService) AddSubjectInClass(ctx context.Context, request *textService.AddSubjectInClassRequest) (*textService.AddSubjectInClassResponse, error) {
	return ts.textHandler.AddSubjectInClass(ctx, request)
}

func (ts *TextService) RemoveSubjectFromClass(ctx context.Context, request *textService.RemoveSubjectFromClassRequest) (*textService.RemoveSubjectFromClassResponse, error) {
	return ts.textHandler.RemoveSubjectFromClass(ctx, request)
}

// // Delete

func (ts *TextService) DeleteClass(ctx context.Context, request *textService.DeleteClassRequest) (*textService.DeleteClassResponse, error) {
	return ts.textHandler.DeleteClass(ctx, request)
}

//-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
// Subject

func (ts *TextService) CreateSubject(ctx context.Context, request *textService.CreateSubjectRequest) (*textService.CreateSubjectResponse, error) {
	return ts.textHandler.CreateSubject(ctx, request)
}

func (ts *TextService) GetSubject(ctx context.Context, request *textService.GetSubjectRequest) (*textService.GetSubjectResponse, error) {
	return ts.textHandler.GetSubject(ctx, request)
}

func (ts *TextService) GetSubjects(ctx context.Context, request *textService.GetSubjectsRequest) (*textService.GetSubjectsResponse, error) {
	return ts.textHandler.GetSubjects(ctx)
}

func (ts *TextService) AssignSectionToSubject(ctx context.Context, request *textService.AssignSectionToSubjectRequest) (*textService.AssignSectionToSubjectResponse, error) {
	return ts.textHandler.AssignSectionToSubject(ctx, request)
}

func (ts *TextService) DeleteSubject(ctx context.Context, request *textService.DeleteSubjectRequest) (*textService.DeleteSubjectResponse, error) {
	return ts.textHandler.DeleteSubject(ctx, request)
}

//-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
// Section

func (ts *TextService) CreateSection(ctx context.Context, request *textService.CreateSectionRequest) (*textService.CreateSectionResponse, error) {
	return ts.textHandler.CreateSection(ctx, request)
}

func (ts *TextService) GetSection(ctx context.Context, request *textService.GetSectionRequest) (*textService.GetSectionResponse, error) {
	return ts.textHandler.GetSection(ctx, request)
}

func (ts *TextService) GetSections(ctx context.Context, request *textService.GetSectionsRequest) (*textService.GetSectionsResponse, error) {
	return ts.textHandler.GetSections(ctx)
}

func (ts *TextService) AddLessonInSection(ctx context.Context, request *textService.AddLessonInSectionRequest) (*textService.AddLessonInSectionResponse, error) {
	return ts.textHandler.AddLessonInSection(ctx, request)
}

func (ts *TextService) RemoveLessonFromSection(ctx context.Context, request *textService.RemoveLessonFromSectionRequest) (*textService.RemoveLessonFromSectionResponse, error) {
	return ts.textHandler.RemoveLessonFromSection(ctx, request)
}

func (ts *TextService) DeleteSection(ctx context.Context, request *textService.DeleteSectionRequest) (*textService.DeleteSectionResponse, error) {
	return ts.textHandler.DeleteSection(ctx, request)
}

//-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
// Lesson

func (ts *TextService) CreateLesson(ctx context.Context, request *textService.CreateLessonRequest) (*textService.CreateLessonResponse, error) {
	return ts.textHandler.CreateLesson(ctx, request)
}

func (ts *TextService) GetLesson(ctx context.Context, request *textService.GetLessonRequest) (*textService.GetLessonResponse, error) {
	return ts.textHandler.GetLesson(ctx, request)
}

func (ts *TextService) GetLessons(ctx context.Context, request *textService.GetLessonsRequest) (*textService.GetLessonsResponse, error) {
	return ts.textHandler.GetLessons(ctx, request)
}

func (ts *TextService) AddVideoInLesson(ctx context.Context, request *textService.AddVideoInLessonRequest) (*textService.AddVideoInLessonResponse, error) {
	return ts.textHandler.AddVideoInLesson(ctx, request)
}

func (ts *TextService) RemoveVideoFromLesson(ctx context.Context, request *textService.RemoveVideoFromLessonRequest) (*textService.RemoveVideoFromLessonResponse, error) {
	return ts.textHandler.RemoveVideoFromLesson(ctx, request)
}

func (ts *TextService) AddFileInLesson(ctx context.Context, request *textService.AddFileInLessonRequest) (*textService.AddFileInLessonResponse, error) {
	return ts.textHandler.AddFileInLesson(ctx, request)
}

func (ts *TextService) RemoveFileFromLesson(ctx context.Context, request *textService.RemoveFileFromLessonRequest) (*textService.RemoveFileFromLessonResponse, error) {
	return ts.textHandler.RemoveFileFromLesson(ctx, request)
}

func (ts *TextService) AddExerciseInLesson(ctx context.Context, request *textService.AddExerciseInLessonRequest) (*textService.AddExerciseInLessonResponse, error) {
	return ts.textHandler.AddExerciseInLesson(ctx, request)
}

func (ts *TextService) RemoveExerciseFromLesson(ctx context.Context, request *textService.RemoveExerciseFromLessonRequest) (*textService.RemoveExerciseFromLessonResponse, error) {
	return ts.textHandler.RemoveExerciseFromLesson(ctx, request)
}

func (ts *TextService) AddCommentInLesson(ctx context.Context, request *textService.AddCommentInLessonRequest) (*textService.AddCommentInLessonResponse, error) {
	return ts.textHandler.AddCommentInLesson(ctx, request)
}

func (ts *TextService) RemoveCommentFromLesson(ctx context.Context, request *textService.RemoveCommentFromLessonRequest) (*textService.RemoveCommentFromLessonResponse, error) {
	return ts.textHandler.RemoveCommentFromLesson(ctx, request)
}

func (ts *TextService) IncreaseRating(ctx context.Context, request *textService.IncreaseRatingRequest) (*textService.IncreaseRatingResponse, error) {
	return ts.textHandler.IncreaseRating(ctx, request)
}

func (ts *TextService) DeleteLesson(ctx context.Context, request *textService.DeleteLessonRequest) (*textService.DeleteLessonResponse, error) {
	return ts.textHandler.DeleteLesson(ctx, request)
}
