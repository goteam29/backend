package repository

import (
	textService "api-repository/pkg/api/text-service"
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func InsertClass(ctx context.Context, db *sql.DB, req *textService.CreateClassRequest) (*textService.CreateClassResponse, error) {
	id := uuid.New()

	_, err := db.ExecContext(ctx, "INSERT INTO classes (id, number) VALUES ($1, $2)", id, req.Number)
	if err != nil {
		return nil, fmt.Errorf("pgInsertClass: failed to insert class in classes: %w", err)
	}

	return &textService.CreateClassResponse{
		Id: id.String(),
	}, nil
}

func SelectClass(ctx context.Context, db *sql.DB, req *textService.GetClassRequest) (*textService.GetClassResponse, error) {
	var (
		id         string
		number     int32
		subjectIds pq.StringArray
	)

	tx, err := db.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, fmt.Errorf("pgSelectClass: failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	classRow := tx.QueryRowContext(ctx, "SELECT id, number FROM classes WHERE id = $1", req.Id)
	err = classRow.Scan(&id, &number)
	if err != nil {
		if err == sql.ErrNoRows {
			tx.Commit()
			return nil, nil
		}
		return nil, fmt.Errorf("pgSelectClass: failed to scan class: %w", err)
	}

	subjectRows, err := tx.QueryContext(ctx, "SELECT subject_id FROM classes_subjects WHERE class_id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("pgSelectClass: failed to select subjects: %w", err)
	}
	defer subjectRows.Close()

	for subjectRows.Next() {
		var subjectId uuid.UUID
		err := subjectRows.Scan(&subjectId)
		if err != nil {
			return nil, fmt.Errorf("pgSelectClass: failed to scan subject id: %w", err)
		}
		subjectIds = append(subjectIds, subjectId.String())
	}

	if err := subjectRows.Err(); err != nil {
		return nil, fmt.Errorf("pgSelectClass: failed to iterate over subject rows: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("pgSelectClass: failed to commit transaction: %w", err)
	}

	return &textService.GetClassResponse{
		Class: &textService.Class{
			Id:         id,
			Number:     number,
			SubjectIds: subjectIds,
		},
	}, nil
}

func SelectClasses(ctx context.Context, db *sql.DB) (*textService.GetClassesResponse, error) {
	classes := make([]*textService.Class, 0, 11)

	tx, err := db.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, fmt.Errorf("pgSelectClasses: failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	classRows, err := tx.QueryContext(ctx, "SELECT id, number FROM classes")
	if err != nil {
		return nil, fmt.Errorf("pgSelectClasses: failed to select classes from database: %w", err)
	}
	defer classRows.Close()

	for classRows.Next() {
		var (
			id         string
			number     int32
			subjectIds pq.StringArray
		)
		class := &textService.Class{}

		err := classRows.Scan(&id, &number)
		if err != nil {
			return nil, fmt.Errorf("pgSelectClasses: failed to scan rows: %w", err)
		}

		subjectRows, err := tx.QueryContext(ctx, "SELECT subject_id FROM classes_subjects WHERE class_id = $1", id)
		if err != nil {
			return nil, fmt.Errorf("pgSelectClasses: failed to select subjects: %w", err)
		}
		defer subjectRows.Close()

		for subjectRows.Next() {
			var subjectId uuid.UUID

			err := subjectRows.Scan(&subjectId)
			if err != nil {
				return nil, fmt.Errorf("pgSelectClasses: failed to scan subject id: %w", err)
			}

			subjectIds = append(subjectIds, subjectId.String())
		}

		if err := subjectRows.Err(); err != nil {
			return nil, fmt.Errorf("pgSelectClasses: failed to iterate over subject rows: %w", err)
		}

		class.Id = id
		class.Number = number
		class.SubjectIds = subjectIds

		classes = append(classes, class)
	}

	if err := classRows.Err(); err != nil {
		return nil, fmt.Errorf("pgSelectClasses: failed to iterate over class rows: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("pgSelectClass: failed to commit transaction: %w", err)
	}

	return &textService.GetClassesResponse{
		Classes: classes,
	}, nil
}

func AddSubjectInClass(ctx context.Context, db *sql.DB, req *textService.AddSubjectInClassRequest) (*textService.AddSubjectInClassResponse, error) {
	var subjectId string

	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("pgAddSubjectInClass: failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	subjectRow := tx.QueryRowContext(ctx, "SELECT id FROM subjects WHERE id = $1", req.SubjectId)
	err = subjectRow.Scan(&subjectId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		} else {
			return nil, fmt.Errorf("pgAddSubjectInClass: failed to scan subject: %w", err)
		}
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO classes_subjects (class_id, subject_id) VALUES ($1, $2)", req.Id, req.SubjectId)
	if err != nil {
		return nil, fmt.Errorf("pgAddSubjectInClass: failed to insert subject in classes_subjects: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("pgAddSubjectInClass: failed to commit transaction: %w", err)
	}

	return &textService.AddSubjectInClassResponse{
		SubjectId: req.SubjectId,
	}, nil
}

func RemoveSubjectFromClass(ctx context.Context, db *sql.DB, req *textService.RemoveSubjectFromClassRequest) (*textService.RemoveSubjectFromClassResponse, error) {
	_, err := db.ExecContext(ctx, "DELETE FROM classes_subjects WHERE class_id = $1 AND subject_id = $2", req.Id, req.SubjectId)
	if err != nil {
		return nil, fmt.Errorf("pgUpdateClass: failed to remove subject from classes_subjects: %w", err)
	}

	return &textService.RemoveSubjectFromClassResponse{
		SubjectId: req.SubjectId,
	}, nil
}

func DeleteClass(ctx context.Context, db *sql.DB, req *textService.DeleteClassRequest) (*textService.DeleteClassResponse, error) {
	_, err := db.ExecContext(ctx, "DELETE FROM classes WHERE id = $1", req.Id)
	if err != nil {
		return nil, fmt.Errorf("pgDeleteClass: failed to delete class in database: %w", err)
	}

	return &textService.DeleteClassResponse{
		Id: req.Id,
	}, nil
}
