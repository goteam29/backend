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

// func SelectSubject(db *sql.DB, req *textService.GetSubjectRequest) (*textService.GetSubjectResponse, error) {
// 	subjectResponse := &textService.GetSubjectResponse{
// 		Subject: &textService.Subject{
// 			SectionIds: make([]string, 0),
// 		},
// 	}

// 	subject := db.QueryRow("SELECT number, subject_ids FROM subjects WHERE number = ($1)", req.Number)
// 	err := subject.Scan(&subjectResponse.Subject.Number, pq.Array(&subjectResponse.Subject.SubjectIds))
// 	if err != nil {
// 		return nil, fmt.Errorf("pgSelectSubject: failed to scan subject: %v", err)
// 	}

// 	return subjectResponse, nil
// }

// func SelectSubjects(db *sql.DB) (*textService.GetSubjectesResponse, error) {
// 	subjects, err := db.Query("SELECT number, subject_ids FROM subjects")
// 	if err != nil {
// 		return nil, fmt.Errorf("pgSelectSubjectes: failed to select subjects from database: %v", err)
// 	}
// 	defer subjects.Close()

// 	subjectsResponse := make([]*textService.Subject, 0, 11)

// 	for subjects.Next() {
// 		subject := &textService.Subject{}
// 		var subjectIds pq.StringArray

// 		err := subjects.Scan(&subject.Number, &subjectIds)
// 		if err != nil {
// 			return nil, fmt.Errorf("pgSelectSubjectes: failed to scan rows: %v", err)
// 		}

// 		subject.SubjectIds = subjectIds

// 		subjectsResponse = append(subjectsResponse, subject)
// 	}

// 	if err := subjects.Err(); err != nil {
// 		return nil, fmt.Errorf("pgSelectSubjectes: error during rows iteration: %v", err)
// 	}

// 	return &textService.GetSubjectesResponse{
// 		Subjectes: subjectsResponse,
// 	}, nil
// }
