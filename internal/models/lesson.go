package models

import "github.com/google/uuid"

type Lesson struct {
	Id          uuid.UUID
	SectionID   uuid.UUID
	Name        string
	Description string
	VideoIDs    []uuid.UUID
	FileIDs     []uuid.UUID
	ExerciseIDs []uuid.UUID
	CommentIDs  []uuid.UUID
	Rating      int
}
