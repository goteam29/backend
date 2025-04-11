package models

import "github.com/google/uuid"

type Subject struct {
	Id          uuid.UUID
	Name        string
	ClassNumber int
	Sections    []uuid.UUID
}
