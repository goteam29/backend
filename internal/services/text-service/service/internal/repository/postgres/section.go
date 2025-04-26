package repository

import (
	textService "api-repository/pkg/api/text-service"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"google.golang.org/protobuf/types/known/emptypb"
)

func InsertSection(db *sql.DB, req *textService.CreateSectionRequest) (*textService.CreateSectionResponse, error) {
	id := uuid.New()

	_, err := db.Exec("INSERT INTO sections (id, subject_id, name, description) VALUES ($1, $2, $3, $4)",
		id,
		req.SubjectId,
		req.Name,
		req.Description,
	)
	if err != nil {
		return nil, fmt.Errorf("pgInsertSection: failed to insert section into database: %v", err)
	}

	return &textService.CreateSectionResponse{
		Id: id.String(),
	}, nil
}

func SelectSection(db *sql.DB, req *textService.GetSectionRequest) (*textService.GetSectionResponse, error) {
	query := `
		SELECT 
			s.id AS section_id,
			s.subject_id AS subject_id,
			s.name AS section_name,
			s.description AS section_description,
			array_agg(DISTINCT l.id) FILTER (WHERE l.id IS NOT NULL) AS lesson_ids
		FROM
			sections s
		LEFT JOIN
			lessons l on s.id = l.section_id
		WHERE
			s.id = $1
		GROUP BY
			s.id, s.subject_id, s.name, s.description;
	`

	sectionRow := db.QueryRow(query, req.Id)

	var (
		id, subjectId, name, description string
		lessonIds                        pq.StringArray
	)

	err := sectionRow.Scan(&id, &subjectId, &name, &description, &lessonIds)
	if err != nil {
		return nil, fmt.Errorf("pgSelectSection: failed to scan section: %v", err)
	}

	section := &textService.Section{
		Id:          id,
		SubjectId:   subjectId,
		Name:        name,
		Description: description,
		LessonIds:   lessonIds,
	}

	return &textService.GetSectionResponse{
		Section: section,
	}, nil
}

func SelectSections(db *sql.DB) (*textService.GetSectionsResponse, error) {
	sections := make([]*textService.Section, 0, 10)

	query := `
		SELECT 
			s.id AS section_id,
			s.subject_id AS subject_id,
			s.name AS section_name,
			s.description AS section_description,
			array_agg(DISTINCT l.id) FILTER (WHERE l.id IS NOT NULL) AS lesson_ids
		FROM
			sections s
		LEFT JOIN
			lessons l on s.id = l.section_id
		GROUP BY
			s.id, s.subject_id, s.name, s.description;
	`

	sectionRows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("pgSelectSections: failed to query sections: %v", err)
	}
	defer sectionRows.Close()

	for sectionRows.Next() {
		var (
			id, subjectId, name, description string
			lessonIds                        pq.StringArray
		)

		err := sectionRows.Scan(&id, &subjectId, &name, &description, &lessonIds)
		if err != nil {
			return nil, fmt.Errorf("pgSelectSections: failed to scan section: %v", err)
		}

		section := &textService.Section{
			Id:          id,
			SubjectId:   subjectId,
			Name:        name,
			Description: description,
			LessonIds:   lessonIds,
		}

		sections = append(sections, section)
	}

	if err := sectionRows.Err(); err != nil {
		return nil, fmt.Errorf("pgSelectSections: failed to iterate over sections: %v", err)
	}

	return &textService.GetSectionsResponse{
		Sections: sections,
	}, nil
}

func UpdateSection(db *sql.DB, req *textService.AssignLessonToSectionRequest) (*textService.AssignLessonToSectionResponse, error) {
	query := `
		UPDATE lessons
		SET section_id = $1
		WHERE id = $2;
	`

	_, err := db.Exec(query, req.Id, req.LessonId)
	if err != nil {
		return nil, fmt.Errorf("pgAddLessonsInSection: failed to add lesson in section: %v", err)
	}

	return &textService.AssignLessonToSectionResponse{
		LessonId: req.LessonId,
	}, nil
}

func DeleteSection(db *sql.DB, req *textService.DeleteSectionRequest) (*emptypb.Empty, error) {
	_, err := db.Exec("DELETE FROM sections WHERE id = $1", req.Id)
	if err != nil {
		return nil, fmt.Errorf("pgDeleteSection: failed to delete section: %v", err)
	}

	return &emptypb.Empty{}, nil
}
