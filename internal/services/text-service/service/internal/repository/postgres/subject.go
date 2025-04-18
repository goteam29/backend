package repository

import (
	textService "api-repository/pkg/api/text-service"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func InsertSubject(db *sql.DB, req *textService.CreateSubjectRequest, id uuid.UUID) error {
	_, err := db.Exec("INSERT INTO subjects (id, name, class_number, section_ids) VALUES ($1, $2, $3, $4)",
		id, req.Subject.Name, req.Subject.ClassNumber, pq.Array(req.Subject.SectionIds))
	if err != nil {
		return fmt.Errorf("pgInsertSubject: failed to insert subject into database: %v", err)
	}

	return nil
}

func SelectSubject(db *sql.DB, req *textService.GetSubjectRequest) (*textService.GetSubjectResponse, error) {
	subjectResponse := &textService.GetSubjectResponse{
		Subject: &textService.Subject{
			Id:          "",
			Name:        "",
			ClassNumber: 0,
			SectionIds:  make([]string, 0),
		},
	}

	subject := db.QueryRow("SELECT id, name, class_number, section_ids FROM subjects WHERE id = ($1)", req.Id)
	err := subject.Scan(&subjectResponse.Subject.Id, &subjectResponse.Subject.Name, &subjectResponse.Subject.ClassNumber, pq.Array(&subjectResponse.Subject.SectionIds))
	if err != nil {
		return nil, fmt.Errorf("pgSelectSubject: failed to scan subject: %v", err)
	}

	return subjectResponse, nil
}

func SelectSubjects(db *sql.DB) (*textService.GetSubjectsResponse, error) {
	subjects, err := db.Query("SELECT id, name, class_number, section_ids FROM subjects")
	if err != nil {
		return nil, fmt.Errorf("pgSelectSubjects: failed to select subjects from database: %v", err)
	}
	defer subjects.Close()

	subjectsResponse := make([]*textService.Subject, 0, 11)

	for subjects.Next() {
		subject := &textService.Subject{}
		var sectionIds pq.StringArray

		err := subjects.Scan(&subject.Id, &subject.Name, &subject.ClassNumber, &sectionIds)
		if err != nil {
			return nil, fmt.Errorf("pgSelectSubjectes: failed to scan rows: %v", err)
		}

		subject.SectionIds = sectionIds

		subjectsResponse = append(subjectsResponse, subject)
	}

	if err := subjects.Err(); err != nil {
		return nil, fmt.Errorf("pgSelectSubjectes: error during rows iteration: %v", err)
	}

	return &textService.GetSubjectsResponse{
		Subjects: subjectsResponse,
	}, nil
}

func UpdateSubject(db *sql.DB, req *textService.UpdateSubjectRequest) error {	
	_, err := db.Exec("UPDATE subjects SET name = $1, class_number = $2, section_ids = array_append(section_ids, $3) WHERE id = $4",
		req.Subject.Name, req.Subject.ClassNumber, req.Subject.SectionIds[0], req.Subject.Id)
	if err != nil {
		return fmt.Errorf("pgUpdateSubject: failed to update subject in database: %v", err)
	}

	return nil
}
