package repository

import (
	textService "api-repository/pkg/api/text-service"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func InsertLesson(id uuid.UUID, pg *sql.DB, req *textService.CreateLessonRequest) error {
	sectionID := uuid.New()

	_, err := pg.Exec("INSERT INTO lessons (id, section_id, name, description) VALUES ($1, $2, $3, $4)",
		id,
		sectionID,
		req.Lesson.Name,
		req.Lesson.Description,
	)
	if err != nil {
		return fmt.Errorf("pgInsertLesson: failed to insert lesson into database: %v", err)
	}

	return nil
}

func SelectLesson(pg *sql.DB, req *textService.GetLessonRequest) (*textService.GetLessonResponse, error) {
	lessonResponse := &textService.GetLessonResponse{}

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
