package repository

import (
	textService "api-repository/pkg/api/text-service"
	"database/sql"
	"fmt"
	"strings"
)

func InsertClass(db *sql.DB, req *textService.CreateClassRequest) error {
	_, err := db.Exec("INSERT INTO classes (number) VALUES ($1)", req.Class.Number)
	if err != nil {
		return fmt.Errorf("pgInsert: failed to insert class into database: %v", err)
	}

	return nil
}

func SelectClasses(db *sql.DB) (*textService.GetClassesResponse, error) {
	classes, err := db.Query("SELECT * FROM classes")
	if err != nil {
		return nil, fmt.Errorf("pgSelect: failed to select classes from database: %v", err)
	}

	classesResponse := make([]*textService.Class, 0, 11)

	var (
		subjectIdsNullable sql.NullString
	)

	for classes.Next() {
		class := &textService.Class{}
		
		err := classes.Scan(&class.Number, &subjectIdsNullable)
		if err != nil {
			return nil, fmt.Errorf("pgSelect: failed to scan rows: %v", err)
		}

		if subjectIdsNullable.Valid {
			class.SubjectIds = strings.Split(subjectIdsNullable.String, ",")
		} else {
			class.SubjectIds = []string{}
		}
		classesResponse = append(classesResponse, class)
	}

	return &textService.GetClassesResponse{
		Class: classesResponse,
	}, nil
}
