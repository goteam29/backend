package handlers

import (
	postgresRepo "api-repository/internal/services/text-service/service/internal/repository/postgres"
	redisRepo "api-repository/internal/services/text-service/service/internal/repository/redis"
	textService "api-repository/pkg/api/text-service"
	"context"
	"fmt"
)

func (th *TextHandler) CreateClass(ctx context.Context, req *textService.CreateClassRequest) (*textService.CreateClassResponse, error) {
	err := redisRepo.AddClass(th.redis, req)
	if err != nil {
		return nil, fmt.Errorf("createClass: %v", err)
	}

	err = postgresRepo.InsertClass(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("createClass: %v", err)
	}

	return &textService.CreateClassResponse{
		Response: "class created successfully",
	}, nil
}

func (th *TextHandler) GetClass(ctx context.Context, req *textService.GetClassRequest) (*textService.GetClassResponse, error) {
	class, err := redisRepo.GetClass(ctx, th.redis, req)
	if err != nil {
		return nil, fmt.Errorf("getClass: %v", err)
	}

	if class == nil {
		class, err = postgresRepo.SelectClass(th.pg, req)
		if err != nil {
			return nil, fmt.Errorf("getClass: %v", err)
		}

		return class, nil
	}

	return class, nil
}

func (th *TextHandler) GetClasses(ctx context.Context) (*textService.GetClassesResponse, error) {
	classes, err := redisRepo.GetClasses(th.redis, ctx)
	if err != nil {
		return nil, fmt.Errorf("getClasses: %v", err)
	}

	if len(classes.Classes) == 0 {
		classes, err := postgresRepo.SelectClasses(th.pg)
		if err != nil {
			return nil, fmt.Errorf("getClasses: %v", err)
		}

		return classes, nil
	}

	return classes, nil
}
