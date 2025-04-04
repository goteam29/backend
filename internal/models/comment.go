package models

import (
	"github.com/google/uuid"
	"time"
)

type Comment struct {
	Id       uuid.UUID
	Username string
	UserId   uuid.UUID
	Text     string
	Date     time.Time
	Rating   int
	Comments []Comment
}
