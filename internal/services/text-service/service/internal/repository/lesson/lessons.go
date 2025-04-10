package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	textService "api-repository/pkg/api/text-service"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func PgInsert(id uuid.UUID, pg *sql.DB, req *textService.CreateLessonRequest) error {
	sectionID := uuid.New()

	_, err := pg.Exec("INSERT INTO lessons (id, section_id, name, description) VALUES ($1, $2, $3, $4)",
		id,
		sectionID,
		req.Lesson.Name,
		req.Lesson.Description,
	)
	if err != nil {
		return fmt.Errorf("pgInsert: failed to insert lesson into database: %v", err)
	}

	return nil
}

func RedisAdd(id uuid.UUID, rds *redis.Client, req *textService.CreateLessonRequest) error {
	lessonJSON, err := json.Marshal(req.Lesson)
	if err != nil {
		return fmt.Errorf("redisAdd: failed to marshal lesson to JSON: %v", err)
	}

	err = rds.Set(context.Background(), id.String(), string(lessonJSON), 0).Err()
	if err != nil {
		return fmt.Errorf("redisAdd: failed to set lesson in Redis: %v", err)
	}

	log.Print("lesson added to Redis: key: ", id, ", value: ", string(lessonJSON))

	return nil
}

func PgSelect(pg *sql.DB, req *textService.GetLessonRequest) (*textService.GetLessonResponse, error) {
	lesson := pg.QueryRow("SELECT id, section_id, name, description FROM lessons WHERE id = $1", req.Id)
	var lessonID, sectionID, name, description string
	if err := lesson.Scan(&lessonID, &sectionID, &name, &description); err != nil {
		return nil, fmt.Errorf("pgSelect: failed to scan lesson: %v", err)
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

func RedisGet(rds *redis.Client, ctx context.Context, id string) (*textService.GetLessonResponse, error) {
	lessonJSON, err := rds.Get(ctx, id).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("redisGet: failed to get lesson from Redis: %v", err)
	}

	var lesson *textService.Lesson
	err = json.Unmarshal([]byte(lessonJSON), &lesson)
	if err != nil {
		return nil, fmt.Errorf("redisGet: failed to unmarshal lesson from JSON: %v", err)
	}

	log.Print("redisGet: lesson: ", lesson)

	return &textService.GetLessonResponse{
		Lesson: lesson,
	}, nil
}
