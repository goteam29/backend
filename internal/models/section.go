package models

import "github.com/google/uuid"

type Section struct {
	Id          uuid.UUID
	SubjectId   uuid.UUID
	Name        string
	Description string
	LessonIDs   []uuid.UUID
}
