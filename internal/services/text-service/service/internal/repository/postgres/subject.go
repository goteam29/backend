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

func AddSectionInSubject(db *sql.DB, req *textService.AddSectionInSubjectRequest) error {
	if len(req.SectionIds) == 0 {
		return fmt.Errorf("pgUpdateSubject: no section IDs provided")
	}

	if len(req.SectionIds) > 0 {
		_, err := db.Exec("UPDATE subjects SET section_ids = section_ids || $1 WHERE id = $2", pq.Array(req.SectionIds), req.SubjectId)
		if err != nil {
			return fmt.Errorf("pgUpdateSubject: failed to add sections in subject: %v", err)
		}
	} else {
		_, err := db.Exec("UPDATE subjects SET sections_ids = array_append(section_ids, $1) WHERE id = $2", req.SectionIds[0], req.SubjectId)
		if err != nil {
			return fmt.Errorf("pgUpdateSubject: failed to add section in subject: %v", err)
		}
	}

	return nil
}

func RemoveSectionFromSubject(db *sql.DB, req *textService.RemoveSectionFromSubjectRequest) error {
	_, err := db.Exec("UPDATE subjects SET section_ids = array_remove(section_ids, $1) WHERE id = $2", req.SectionId, req.SubjectId)
	if err != nil {
		return fmt.Errorf("pgUpdateSubject: failed to remove section in subject: %v", err)
	}

	return nil
}

func DeleteSubject(db *sql.DB, req *textService.DeleteSubjectRequest) error {
	_, err := db.Exec("DELETE FROM subjects WHERE id = $1", req.Id)
	if err != nil {
		return fmt.Errorf("pgDeleteSubject: failed to delete subject from database: %v", err)
	}

	return nil
}
