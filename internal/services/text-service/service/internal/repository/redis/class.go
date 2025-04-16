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

func GetClass(ctx context.Context, rds *redis.Client, req *textService.GetClassRequest) (*textService.GetClassResponse, error) {
	classResponse := &textService.GetClassResponse{}

	classJSON, err := rds.HGet(ctx, "class", fmt.Sprint(req.Number)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, fmt.Errorf("redisGetClass: failed to get class from Redis: %v", err)
	}

	err = json.Unmarshal([]byte(classJSON), classResponse)
	if err != nil {
		return nil, fmt.Errorf("redisGetClass: failed to unmarshall class from JSON: %v", err)
	}

	return classResponse, nil
}

func GetClasses(rds *redis.Client, ctx context.Context) (*textService.GetClassesResponse, error) {
	classesMap, err := rds.HGetAll(ctx, "class").Result()
	if err != nil {
		return nil, fmt.Errorf("redisGetClasses: failed to get classes from Redis: %v", err)
	}

	classesResponse := make([]*textService.Class, 0, len(classesMap))

	for _, classJSON := range classesMap {
		var class textService.Class
		err := json.Unmarshal([]byte(classJSON), &class)
		if err != nil {
			return nil, fmt.Errorf("redisGetClasses: failed to unmarshall class from JSON: %v", err)
		}
		classesResponse = append(classesResponse, &class)
	}

	return &textService.GetClassesResponse{
		Classes: classesResponse,
	}, nil
}
