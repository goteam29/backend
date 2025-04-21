package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	textService "api-repository/pkg/api/text-service"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

const timeToLive = 30 * time.Second

func AddLesson(id uuid.UUID, rds *redis.Client, req *textService.CreateLessonRequest) error {
	lessonJSON, err := json.Marshal(req.Lesson)
	if err != nil {
		return fmt.Errorf("redisAdd: failed to marshal lesson to JSON: %v", err)
	}

	err = rds.Set(context.Background(), id.String(), string(lessonJSON), timeToLive).Err()
	if err != nil {
		return fmt.Errorf("redisAdd: failed to set lesson in Redis: %v", err)
	}

	log.Print("lesson added to Redis: key: ", id, ", value: ", string(lessonJSON))

	return nil
}

func GetLesson(rds *redis.Client, ctx context.Context, id string) (*textService.GetLessonResponse, error) {
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
