package models

import "github.com/google/uuid"

type Lesson struct {
	Id          uuid.UUID
	Name        string
	Description string
	VideoIDs    []uuid.UUID
	FileIDs     []uuid.UUID
	ExerciseIDs []uuid.UUID
	CommentIds  []uuid.UUID
	Rating      int
}
