package repository

import (
	textService "api-repository/pkg/api/text-service"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

func InsertClass(db *sql.DB, req *textService.CreateClassRequest) error {
	_, err := db.Exec("INSERT INTO classes (number, subject_ids) VALUES ($1, $2)", req.Class.Number, pq.Array(req.Class.SubjectIds))
	if err != nil {
		return fmt.Errorf("pgInsertClass: failed to insert class into database: %v", err)
	}

	return nil
}

func SelectClass(db *sql.DB, req *textService.GetClassRequest) (*textService.GetClassResponse, error) {
	classResponse := &textService.GetClassResponse{
		Class: &textService.Class{
			SubjectIds: make([]string, 0),
		},
	}

	class := db.QueryRow("SELECT number, subject_ids FROM classes WHERE number = ($1)", req.Number)
	err := class.Scan(&classResponse.Class.Number, pq.Array(&classResponse.Class.SubjectIds))
	if err != nil {
		return nil, fmt.Errorf("pgSelectClass: failed to scan class: %v", err)
	}

	return classResponse, nil
}

func SelectClasses(db *sql.DB) (*textService.GetClassesResponse, error) {
	classes, err := db.Query("SELECT number, subject_ids FROM classes")
	if err != nil {
		return nil, fmt.Errorf("pgSelectClasses: failed to select classes from database: %v", err)
	}
	defer classes.Close()

	classesResponse := make([]*textService.Class, 0, 11)

	for classes.Next() {
		class := &textService.Class{}
		var subjectIds pq.StringArray

		err := classes.Scan(&class.Number, &subjectIds)
		if err != nil {
			return nil, fmt.Errorf("pgSelectClasses: failed to scan rows: %v", err)
		}

		class.SubjectIds = subjectIds

		classesResponse = append(classesResponse, class)
	}

	if err := classes.Err(); err != nil {
		return nil, fmt.Errorf("pgSelectClasses: error during rows iteration: %v", err)
	}

	return &textService.GetClassesResponse{
		Classes: classesResponse,
	}, nil
}
