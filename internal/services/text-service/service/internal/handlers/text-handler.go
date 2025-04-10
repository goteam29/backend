package handlers

import (
	textService "api-repository/pkg/api/text-service"
	"context"
	"database/sql"
	"log"

	"github.com/google/uuid"
)

type TextHandler struct {
	db *sql.DB
}

func NewTextHandler(pgConn *sql.DB) *TextHandler {
	return &TextHandler{
		db: pgConn,
	}
}

func (th *TextHandler) CreateLesson(ctx context.Context, req *textService.CreateLessonRequest) (*textService.CreateLessonResponse, error) {
	log.Print("CreateLesson called")
	log.Print("Request: ", req)

	id := uuid.New()
	sectionID := uuid.New() // Assuming you want to create a new section ID for the lesson

	_, err := th.db.Exec("INSERT INTO lessons (id, section_id, name, description) VALUES ($1, $2, $3, $4)",
		id,
		sectionID,
		req.Lesson.Name,
		req.Lesson.Description,
	)
	if err != nil {
		log.Printf("Error inserting lesson into database: %v", err)
		return nil, err
	}

	return &textService.CreateLessonResponse{
		Response: "Lesson created successfully",
	}, nil
}

func (th *TextHandler) GetLesson(ctx context.Context, req *textService.GetLessonRequest) (*textService.GetLessonResponse, error) {
	log.Print("GetLesson called")
	log.Print("Request: ", req)

	lesson:= th.db.QueryRow("SELECT id, section_id, name, description FROM lessons WHERE id = $1", req.Id)
	var lessonID, sectionID, name, description string
	if err := lesson.Scan(&lessonID, &sectionID, &name, &description); err != nil {
		log.Printf("Error getting lesson from database: %v", err)
		return nil, err
	}

	return &textService.GetLessonResponse{
		Lesson: &textService.Lesson{
			Id:          lessonID,
			SectionId:   sectionID,
			Name:        name,
			Description: description,
		},
	}, nil
}
