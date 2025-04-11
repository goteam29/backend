package repository

import (
	textService "api-repository/pkg/api/text-service"
	"database/sql"
	"fmt"
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
	var (
		number   int32
		subjects []string
	)

	classResponse:=make([]*textService.Class, 11)

	var i int
	for classes.Next() {
		err := classes.Scan(&number, &subjects[i])
		if err!=nil{
			return nil, fmt.Errorf("pgSelect: failed to scan classes: %v", err)
		}
		i++
	}

	return *textService.GetClassesResponse{
		Class: []*textService.Class{
			Number:     number,
			SubjectIds: subjects,
		},
	}, nil
}
