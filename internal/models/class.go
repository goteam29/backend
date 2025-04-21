package models

import "github.com/google/uuid"

type Class struct {
	Number   int
	Subjects []uuid.UUID
}
