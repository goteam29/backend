package repository

import (
	textService "api-repository/pkg/api/text-service"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

func InsertSection(db *sql.Tx, req *textService.CreateSectionRequest, id uuid.UUID) error {
	_, err := db.Exec("INSERT INTO sections (id, subject_id, name, description) VALUES ($1, $2, $3, $4)", id, req.SubjectId, req.Name, req.Description)
	if err != nil {
		return fmt.Errorf("pgInsertSection: failed to insert section into database: %v", err)
	}

	return nil
}

// func SelectSection(db *sql.DB, req *textService.GetSectionRequest) (*textService.GetSectionResponse, error) {
// 	sectionResponse := &textService.GetSectionResponse{
// 		Section: &textService.Section{
// 			SubjectIds: make([]string, 0),
// 		},
// 	}

// 	section := db.QueryRow("SELECT number, subject_ids FROM sections WHERE number = ($1)", req.SectionNumber)
// 	err := section.Scan(&sectionResponse.Section.Number, pq.Array(&sectionResponse.Section.SubjectIds))
// 	if err != nil {
// 		return nil, fmt.Errorf("pgSelectSection: failed to scan section: %v", err)
// 	}

// 	return sectionResponse, nil
// }

// func SelectSections(db *sql.DB) (*textService.GetSectionsResponse, error) {
// 	sections, err := db.Query("SELECT number, subject_ids FROM sections")
// 	if err != nil {
// 		return nil, fmt.Errorf("pgSelectSections: failed to select sections from database: %v", err)
// 	}
// 	defer sections.Close()

// 	sectionsResponse := make([]*textService.Section, 0, 11)

// 	for sections.Next() {
// 		section := &textService.Section{}
// 		var subjectIds pq.StringArray

// 		err := sections.Scan(&section.Number, &subjectIds)
// 		if err != nil {
// 			return nil, fmt.Errorf("pgSelectSections: failed to scan rows: %v", err)
// 		}

// 		section.SubjectIds = subjectIds

// 		sectionsResponse = append(sectionsResponse, section)
// 	}

// 	if err := sections.Err(); err != nil {
// 		return nil, fmt.Errorf("pgSelectSections: error during rows iteration: %v", err)
// 	}

// 	return &textService.GetSectionsResponse{
// 		Sections: sectionsResponse,
// 	}, nil
// }

// func AddLessonsInSection(db *sql.DB, req *textService.AddLessonsInSectionRequest) error {
// 	_, err := db.Exec("UPDATE sections SET lesson_ids = array_cat(lesson_ids, $1) WHERE number = $2", pq.Array(req.LessonIds), req.SectionNumber)
// 	if err != nil {
// 		return fmt.Errorf("pgAddLessonsInSection: failed to add lessons in section: %v", err)
// 	}

// 	return nil
// }

// func RemoveLessonFromSection(db *sql.DB, req *textService.RemoveLessonFromSectionRequest) error {
// 	_, err := db.Exec("UPDATE sections SET lesson_ids = array_remove(lesson_ids, $1) WHERE number = $2", req.LessonId, req.SectionNumber)
// 	if err != nil {
// 		return fmt.Errorf("pgRemoveLessonsFromSection: failed to remove lessons from section: %v", err)
// 	}

// 	return nil
// }

// func DeleteSection(db *sql.DB, req *textService.DeleteSectionRequest) error {
// 	_, err := db.Exec("DELETE FROM sections WHERE number = $1", req.SectionNumber)
// 	if err != nil {
// 		return fmt.Errorf("pgDeleteSection: failed to delete section: %v", err)
// 	}

// 	return nil
// }
