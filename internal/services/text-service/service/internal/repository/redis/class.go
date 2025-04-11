package repository

import (
	textService "api-repository/pkg/api/text-service"
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func AddClass(rds *redis.Client, req *textService.CreateClassRequest) error {
	classJSON, err := json.Marshal(req.Class)
	if err != nil {
		return fmt.Errorf("redisAdd: failed to marshal class to JSON: %v", err)
	}

	err = rds.HSet(context.Background(), "class", fmt.Sprint(req.Class.Number), string(classJSON)).Err()
	if err != nil {
		return fmt.Errorf("redisAdd: failed to set class in Redis: %v", err)
	}

	return nil
}

func GetClasses(rds *redis.Client, ctx context.Context) (*textService.GetClassesResponse, error) {
	classes, err := rds.HGetAll(ctx, "class").Result()
	if err != nil {
		return nil, fmt.Errorf("redisGet: failed to get class from Redis: %v", err)
	}

	classesResponse := make([]*textService.Class, len(classes))
	var i int
	for key := range classes {
		err := json.Unmarshal([]byte(classes[key]), &classesResponse[i])
		if err != nil {
			return nil, fmt.Errorf("redisGet: failed to unmarshall class from JSON: %v", err)
		}
		i++
	}

	return &textService.GetClassesResponse{
		Class: classesResponse,
	}, nil
}
