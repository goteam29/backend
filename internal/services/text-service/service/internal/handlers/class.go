package handlers

import (
	postgresRepo "api-repository/internal/services/text-service/service/internal/repository/postgres"
	textService "api-repository/pkg/api/text-service"
	"context"
	"fmt"
)

func (th *TextHandler) CreateClass(ctx context.Context, req *textService.CreateClassRequest) (*textService.CreateClassResponse, error) {
	// err := redisRepo.AddClass(th.redis, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("createClass: %v", err)
	// }

	id, err := postgresRepo.InsertClass(ctx, th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("createClass: %v", err)
	}

	return id, nil
}

func (th *TextHandler) GetClass(ctx context.Context, req *textService.GetClassRequest) (*textService.GetClassResponse, error) {
	// class, err := redisRepo.GetClass(ctx, th.redis, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("getClass: %v", err)
	// }

	// if class == nil {
	class, err := postgresRepo.SelectClass(ctx, th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("getClass: %v", err)
	}

	return class, nil
	// }

	// return class, nil
}

func (th *TextHandler) GetClasses(ctx context.Context) (*textService.GetClassesResponse, error) {
	// classes, err := redisRepo.GetClasses(th.redis, ctx)
	// if err != nil {
	// 	return nil, fmt.Errorf("getClasses: %v", err)
	// }

	// if len(classes.Classes) == 0 {
	classes, err := postgresRepo.SelectClasses(ctx, th.pg)
	if err != nil {
		return nil, fmt.Errorf("getClasses: %v", err)
	}

	return classes, nil
	// }

	// return classes, nil
}

func (th *TextHandler) AddSubjectInClass(ctx context.Context, req *textService.AddSubjectInClassRequest) (*textService.AddSubjectInClassResponse, error) {
	subjectId, err := postgresRepo.AddSubjectInClass(ctx, th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("updateClass: %v", err)
	}

	if subjectId == nil {
		return nil, fmt.Errorf("there is no such subject")
	}

	return subjectId, nil
}

func (th *TextHandler) RemoveSubjectFromClass(ctx context.Context, req *textService.RemoveSubjectFromClassRequest) (*textService.RemoveSubjectFromClassResponse, error) {
	subjectId, err := postgresRepo.RemoveSubjectFromClass(ctx, th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("updateClass: %v", err)
	}

	return subjectId, nil
}

func (th *TextHandler) DeleteClass(ctx context.Context, req *textService.DeleteClassRequest) (*textService.DeleteClassResponse, error) {
	// err := redisRepo.DeleteClass(th.redis, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("deleteClass: %v", err)
	// }

	id, err := postgresRepo.DeleteClass(ctx, th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("deleteClass: %v", err)
	}

	return id, nil
}
