package repository

import (
	textService "api-repository/pkg/api/text-service"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"google.golang.org/protobuf/types/known/emptypb"
)

var ErrClassAlreadyExists = errors.New("class already exists")	
var ErrClassNotFound = errors.New("class not found")
var ErrClassOrSubjectNotFound = errors.New("class or subject not found")

func InsertClass(ctx context.Context, db *sql.DB, req *textService.CreateClassRequest) (*textService.CreateClassResponse, error) {
	id := uuid.New()

	_, err := db.ExecContext(ctx, "INSERT INTO classes (id, number) VALUES ($1, $2)", id, req.Number)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return nil, ErrClassAlreadyExists
		}
		return nil, fmt.Errorf("pgInsertClass: failed to insert class in classes: %w", err)
	}

	return &textService.CreateClassResponse{
		Id: id.String(),
	}, nil
}

func SelectClass(ctx context.Context, db *sql.DB, req *textService.GetClassRequest) (*textService.GetClassResponse, error) {
	query := `
		SELECT
			c.id,
			c.number,
			array_agg(cs.subject_id) FILTER (WHERE cs.subject_id IS NOT NULL) AS subject_ids
		FROM
			classes c
		LEFT JOIN
			classes_subjects cs ON c.id = cs.class_id
		WHERE
			c.id = $1
		GROUP BY
			c.id, c.number;
	`

	classRow := db.QueryRowContext(ctx, query, req.Id)

	var (
		id         string
		number     int32
		subjectIDs pq.StringArray
	)

	err := classRow.Scan(&id, &number, &subjectIDs)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrClassNotFound
		}
		return nil, fmt.Errorf("pgSelectClass: failed to scan row: %w", err)
	}

	class := &textService.Class{
		Id:         id,
		Number:     number,
		SubjectIds: subjectIDs,
	}

	return &textService.GetClassResponse{
		Class: class,
	}, nil
}

func SelectClasses(ctx context.Context, db *sql.DB) (*textService.GetClassesResponse, error) {
	classes := make([]*textService.Class, 0, 11)

	query := `
		SELECT
			c.id,
			c.number,
			array_agg(cs.subject_id) FILTER (WHERE cs.subject_id IS NOT NULL) AS subject_ids
		FROM
			classes c
		LEFT JOIN
			classes_subjects cs ON c.id = cs.class_id
		GROUP BY
			c.id, c.number
		ORDER BY
			c.number;
	`

	classRows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("pgSelectClasses: failed to select classes from database: %w", err)
	}
	defer classRows.Close()

	for classRows.Next() {
		var (
			id         string
			number     int32
			subjectIDs pq.StringArray
		)

		err := classRows.Scan(&id, &number, &subjectIDs)
		if err != nil {
			return nil, fmt.Errorf("pgSelectClasses: failed to scan row: %w", err)
		}

		class := &textService.Class{
			Id:         id,
			Number:     number,
			SubjectIds: subjectIDs,
		}

		classes = append(classes, class)
	}

	if err := classRows.Err(); err != nil {
		return nil, fmt.Errorf("pgSelectClasses: failed to iterate over class rows: %w", err)
	}

	return &textService.GetClassesResponse{
		Classes: classes,
	}, nil
}

func AddSubjectInClass(ctx context.Context, db *sql.DB, req *textService.AddSubjectInClassRequest) (*textService.AddSubjectInClassResponse, error) {
	query := `
		INSERT INTO public.classes_subjects (class_id, subject_id)
		SELECT
    		$1,
    		$2
		WHERE EXISTS (
    		SELECT 1
    		FROM public.subjects
    		WHERE id = $2
		);
	`

	result, err := db.ExecContext(ctx, query, req.Id, req.SubjectId)
	if err != nil {
		return nil, fmt.Errorf("pgAddSubjectInClass: failed to insert subject in classes_subjects: %w", err)
	}

	rowsAddected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("pgAddSubjectInClass: failed to get rows affected: %w", err)
	}
	if rowsAddected == 0 {
		return nil, ErrClassOrSubjectNotFound
	}

	return &textService.AddSubjectInClassResponse{
		SubjectId: req.SubjectId,
	}, nil
}

func RemoveSubjectFromClass(ctx context.Context, db *sql.DB, req *textService.RemoveSubjectFromClassRequest) (*textService.RemoveSubjectFromClassResponse, error) {
	query := `
		DELETE 
		FROM classes_subjects 
		WHERE class_id = $1 
		AND subject_id = $2;
	`

	result, err := db.ExecContext(ctx, query, req.Id, req.SubjectId)
	if err != nil {
		return nil, fmt.Errorf("pgUpdateClass: failed to remove subject from classes_subjects: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("pgRemoveSubjectFromClass: failed to get rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return nil, ErrSubjectNotFound
	}

	return &textService.RemoveSubjectFromClassResponse{
		SubjectId: req.SubjectId,
	}, nil
}

func DeleteClass(ctx context.Context, db *sql.DB, req *textService.DeleteClassRequest) (*emptypb.Empty, error) {
	query := `
		DELETE 
		FROM classes 
		WHERE id = $1;
	`
	result, err := db.ExecContext(ctx, query, req.Id)
	if err != nil {
		return nil, fmt.Errorf("pgDeleteClass: failed to delete class in database: %w", err)
	}

	rowsAffected,err:= result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("pgDeleteClass: failed to get rows affected: %w", err)	
	}

	if rowsAffected == 0 {
		return nil, ErrClassNotFound
	}

	return &emptypb.Empty{}, nil
}
