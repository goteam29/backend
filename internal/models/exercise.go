package models

import "github.com/google/uuid"

type Exercise struct {
	Id       uuid.UUID
	Name     string
	Text     string
	Answer   string
	LessonId uuid.UUID
}
