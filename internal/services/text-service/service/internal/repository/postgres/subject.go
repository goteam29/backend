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

	query := `
        WITH inserted_subject AS (
            INSERT INTO public.subjects (id, name)
            VALUES ($1, $2)
            RETURNING id
        )
        INSERT INTO public.classes_subjects (class_id, subject_id)
        SELECT
            $3,
            inserted_subject.id
        FROM inserted_subject
        WHERE EXISTS (
            SELECT 1
            FROM public.classes
            WHERE id = $3
        );
	`

	_, err := db.Exec(query, id, req.Name, req.ClassId)
	if err != nil {
		return nil, fmt.Errorf("pgInsertSubject: failed to insert subject into database: %v", err)
	}

	return &textService.CreateSubjectResponse{
		Id: id.String(),
	}, nil
}

func SelectSubject(ctx context.Context, db *sql.DB, req *textService.GetSubjectRequest) (*textService.GetSubjectResponse, error) {
	query := `
		SELECT
    		s.id AS subject_id,
    		s.name AS subject_name,
    		array_agg(DISTINCT cs.class_id) FILTER (WHERE cs.class_id IS NOT NULL) AS class_ids,
    		array_agg(DISTINCT sec.id) FILTER (WHERE sec.id IS NOT NULL) AS section_ids
		FROM
    		public.subjects s
		LEFT JOIN
    		public.classes_subjects cs ON s.id = cs.subject_id
		LEFT JOIN
    		public.sections sec ON s.id = sec.subject_id
		WHERE
    		s.id = $1
		GROUP BY
    		s.id, s.name;
	`

	subjectRow := db.QueryRow(query, req.Id)

	var (
		id, name, classId string
		sectionIds        pq.StringArray
	)

	err := subjectRow.Scan(&id, &name, &classId, &sectionIds)
	if err != nil {
		return nil, fmt.Errorf("pgSelectSubject: failed to scan row: %v", err)
	}

	subject := &textService.Subject{
		Id:         id,
		Name:       name,
		ClassId:    classId,
		SectionIds: sectionIds,
	}

	return &textService.GetSubjectResponse{
		Subject: subject,
	}, nil
}

func SelectSubjects(ctx context.Context, db *sql.DB) (*textService.GetSubjectsResponse, error) {
	subjects := make([]*textService.Subject, 0, 10)

	query := `
		SELECT
    		s.id AS subject_id,
    		s.name AS subject_name,
    		array_agg(DISTINCT cs.class_id) FILTER (WHERE cs.class_id IS NOT NULL) AS class_ids,
    		array_agg(DISTINCT sec.id) FILTER (WHERE sec.id IS NOT NULL) AS section_ids
		FROM
    		public.subjects s
		LEFT JOIN
    		public.classes_subjects cs ON s.id = cs.subject_id
		LEFT JOIN
    		public.sections sec ON s.id = sec.subject_id
		GROUP BY
    		s.id, s.name;
	`

	subjectRows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("pgSelectSubjects: failed to query subjects: %v", err)
	}
	defer subjectRows.Close()

	for subjectRows.Next() {
		var (
			id, name, classId string
			sectionIds        pq.StringArray
		)

		err := subjectRows.Scan(&id, &name, &classId, &sectionIds)
		if err != nil {
			return nil, fmt.Errorf("pgSelectSubjects: failed to scan row: %v", err)
		}

		subject := &textService.Subject{
			Id:         id,
			Name:       name,
			ClassId:    classId,
			SectionIds: sectionIds,
		}

		subjects = append(subjects, subject)
	}

	if err := subjectRows.Err(); err != nil {
		return nil, fmt.Errorf("pgSelectSubjects: error iterating over rows: %v", err)
	}

	return &textService.GetSubjectsResponse{
		Subjects: subjects,
	}, nil
}

//func UpdateSubject(ctx context.Context, db *sql.DB, req *textService.AssignSectionToSubjectRequest) (*textService.AssignSectionToSubjectResponse, error) {
//	query := `
//		UPDATE sections
//		SET subject_id = $1
//		WHERE id = $2;
//	`
//
//	_, err := db.Exec(query, req.Id, req.SectionId)
//	if err != nil {
//		return nil, fmt.Errorf("pgUpdateSubject: failed to add section in subject: %v", err)
//	}
//
//	return &textService.AssignSectionToSubjectResponse{
//		SectionId: req.SectionId,
//	}, nil
//}
//
//func DeleteSubject(ctx context.Context, db *sql.DB, req *textService.DeleteSubjectRequest) (*textService.DeleteSubjectResponse, error) {
//	_, err := db.Exec("DELETE FROM subjects WHERE id = $1", req.Id)
//	if err != nil {
//		return nil, fmt.Errorf("pgDeleteSubject: failed to delete subject from database: %v", err)
//	}
//
//	return &textService.DeleteSubjectResponse{
//		Id: req.Id,
//	}, nil
//}
