package repository

import (
	textService "api-repository/pkg/api/text-service"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func InsertSection(db *sql.Tx, req *textService.CreateSectionRequest, id uuid.UUID) error {
	_, err := db.Exec("INSERT INTO sections (id, subject_id, name, description) VALUES ($1, $2, $3, $4)", id, req.SubjectId, req.Name, req.Description)
	if err != nil {
		return fmt.Errorf("pgInsertSection: failed to insert section into database: %v", err)
	}

	return nil
}

func SelectSection(db *sql.DB, req *textService.GetSectionRequest) (*textService.GetSectionResponse, error) {
	sectionResponse := &textService.GetSectionResponse{
		Section: &textService.Section{
			LessonIds: make([]string, 0),
		},
	}

	section := db.QueryRow("SELECT id, subject_id, name, description, lesson_ids FROM sections WHERE id = ($1)", req.Id)
	err := section.Scan(&sectionResponse.Section.Id, &sectionResponse.Section.SubjectId,
		&sectionResponse.Section.Name, &sectionResponse.Section.Description, pq.Array(&sectionResponse.Section.LessonIds))
	if err != nil {
		return nil, fmt.Errorf("pgSelectSection: failed to scan section: %v", err)
	}

	return sectionResponse, nil
}

func SelectSections(db *sql.DB) (*textService.GetSectionsResponse, error) {
	sections, err := db.Query("SELECT id, subject_id, name, description, lesson_ids FROM sections")
	if err != nil {
		return nil, fmt.Errorf("pgSelectSections: failed to select sections from database: %v", err)
	}
	defer sections.Close()

	sectionsResponse := make([]*textService.Section, 0, 5)

	for sections.Next() {
		section := &textService.Section{}
		var lessonIds pq.StringArray

		err := sections.Scan(&section.Id, &section.SubjectId, &section.Name, &section.Description, &lessonIds)
		if err != nil {
			return nil, fmt.Errorf("pgSelectSections: failed to scan rows: %v", err)
		}

		section.LessonIds = lessonIds

		sectionsResponse = append(sectionsResponse, section)
	}

	if err := sections.Err(); err != nil {
		return nil, fmt.Errorf("pgSelectSections: error during rows iteration: %v", err)
	}

	return &textService.GetSectionsResponse{
		Sections: sectionsResponse,
	}, nil
}

func AddLessonInSection(db *sql.DB, req *textService.AddLessonInSectionRequest) (*textService.AddLessonInSectionResponse, error) {
	_, err := db.Exec("UPDATE sections SET lesson_ids = array_append(lesson_ids, $1) WHERE id = $2", req.LessonId, req.Id)
	if err != nil {
		return nil, fmt.Errorf("pgAddLessonsInSection: failed to add lesson in section: %v", err)
	}

	return &textService.AddLessonInSectionResponse{
		LessonId: req.LessonId,
	}, nil
}

func RemoveLessonFromSection(db *sql.DB, req *textService.RemoveLessonFromSectionRequest) (*textService.RemoveLessonFromSectionResponse, error) {
	_, err := db.Exec("UPDATE sections SET lesson_ids = array_remove(lesson_ids, $1) WHERE id = $2", req.LessonId, req.Id)
	if err != nil {
		return nil, fmt.Errorf("pgRemoveLessonsFromSection: failed to remove lessons from section: %v", err)
	}

	return &textService.RemoveLessonFromSectionResponse{
		LessonId: req.LessonId,
	}, nil
}

func DeleteSection(db *sql.DB, req *textService.DeleteSectionRequest) (*textService.DeleteSectionResponse, error) {
	_, err := db.Exec("DELETE FROM sections WHERE id = $1", req.Id)
	if err != nil {
		return nil, fmt.Errorf("pgDeleteSection: failed to delete section: %v", err)
	}

	return &textService.DeleteSectionResponse{
		Id: req.Id,
	}, nil
}
