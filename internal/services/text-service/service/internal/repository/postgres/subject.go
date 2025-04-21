package repository

import (
	textService "api-repository/pkg/api/text-service"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func InsertSubject(ctx context.Context, db *sql.DB, req *textService.CreateSubjectRequest) (*textService.CreateSubjectResponse, error) {
	id := uuid.New()

	_, err := db.Exec("INSERT INTO subjects (id, name) VALUES ($1, $2)",
		id, req.Name)
	if err != nil {
		return nil, fmt.Errorf("pgInsertSubject: failed to insert subject into database: %v", err)
	}

	return &textService.CreateSubjectResponse{
		Id: id.String(),
	}, nil
}

func SelectSubject(ctx context.Context, db *sql.DB, req *textService.GetSubjectRequest) (*textService.GetSubjectResponse, error) {
	subjectResponse := &textService.GetSubjectResponse{
		Subject: &textService.Subject{
			SectionIds: make([]string, 0),
		},
	}

	subject := db.QueryRow("SELECT id, name, class_id, section_ids FROM subjects WHERE id = ($1)", req.Id)
	err := subject.Scan(&subjectResponse.Subject.Id, &subjectResponse.Subject.Name, &subjectResponse.Subject.ClassId, pq.Array(&subjectResponse.Subject.SectionIds))
	if err != nil {
		return nil, fmt.Errorf("pgSelectSubject: failed to scan subject: %v", err)
	}

	return subjectResponse, nil
}

func SelectSubjects(ctx context.Context, db *sql.DB) (*textService.GetSubjectsResponse, error) {
	subjects, err := db.Query("SELECT id, name, class_id, section_ids FROM subjects")
	if err != nil {
		return nil, fmt.Errorf("pgSelectSubjects: failed to select subjects from database: %v", err)
	}
	defer subjects.Close()

	subjectsResponse := make([]*textService.Subject, 0, 11)

	for subjects.Next() {
		subject := &textService.Subject{}
		var sectionIds pq.StringArray

		err := subjects.Scan(&subject.Id, &subject.Name, &subject.ClassId, &sectionIds)
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

func AddSectionInSubject(ctx context.Context, db *sql.Tx, req *textService.AddSectionInSubjectRequest) (*textService.AddSectionInSubjectResponse, error) {
	_, err := db.Exec("UPDATE subjects SET section_ids = array_append(section_ids, $1) WHERE id = $2", req.SectionId, req.Id)
	if err != nil {
		return nil, fmt.Errorf("pgUpdateSubject: failed to add section in subject: %v", err)
	}

	return &textService.AddSectionInSubjectResponse{
		SectionId: req.SectionId,
	}, nil
}

func RemoveSectionFromSubject(ctx context.Context, db *sql.DB, req *textService.RemoveSectionFromSubjectRequest) (*textService.RemoveSectionFromSubjectResponse, error) {
	_, err := db.Exec("UPDATE subjects SET section_ids = array_remove(section_ids, $1) WHERE id = $2", req.SectionId, req.Id)
	if err != nil {
		return nil, fmt.Errorf("pgUpdateSubject: failed to remove section in subject: %v", err)
	}

	return &textService.RemoveSectionFromSubjectResponse{
		SectionId: req.SectionId,
	}, nil
}

func DeleteSubject(ctx context.Context, db *sql.DB, req *textService.DeleteSubjectRequest) (*textService.DeleteSubjectResponse, error) {
	_, err := db.Exec("DELETE FROM subjects WHERE id = $1", req.Id)
	if err != nil {
		return nil, fmt.Errorf("pgDeleteSubject: failed to delete subject from database: %v", err)
	}

	return &textService.DeleteSubjectResponse{
		Id: req.Id,
	}, nil
}
