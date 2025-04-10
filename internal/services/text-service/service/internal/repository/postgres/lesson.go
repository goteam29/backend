package repository

import (
	textService "api-repository/pkg/api/text-service"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func LessonInsert(id uuid.UUID, pg *sql.DB, req *textService.CreateLessonRequest) error {
	sectionID := uuid.New()

	_, err := pg.Exec("INSERT INTO lessons (id, section_id, name, description) VALUES ($1, $2, $3, $4)",
		id,
		sectionID,
		req.Lesson.Name,
		req.Lesson.Description,
	)
	if err != nil {
		return fmt.Errorf("pgInsert: failed to insert lesson into database: %v", err)
	}

	return nil
}

func LessonSelect(pg *sql.DB, req *textService.GetLessonRequest) (*textService.GetLessonResponse, error) {
	lesson := pg.QueryRow("SELECT id, section_id, name, description FROM lessons WHERE id = $1", req.Id)
	var lessonID, sectionID, name, description string
	if err := lesson.Scan(&lessonID, &sectionID, &name, &description); err != nil {
		return nil, fmt.Errorf("pgSelect: failed to scan lesson: %v", err)
	}

	return &textService.GetLessonResponse{
		Lesson: &textService.Lesson{
			Id:          lessonID,
			SectionId:   sectionID,
			Name:        name,
			Description: description,
		},
	}, nil
}
