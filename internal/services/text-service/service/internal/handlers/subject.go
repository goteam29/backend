package handlers

import (
	postgresRepo "api-repository/internal/services/text-service/service/internal/repository/postgres"
	textService "api-repository/pkg/api/text-service"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (th *TextHandler) CreateSubject(ctx context.Context, req *textService.CreateSubjectRequest) (*textService.CreateSubjectResponse, error) {
	// err := redisRepo.AddSubject(th.redis, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("createSubject: %v", err)
	// }

	id := uuid.New()

	err := postgresRepo.InsertSubject(th.pg, req, id)
	if err != nil {
		return nil, fmt.Errorf("createSubject: %v", err)
	}

	updateRequest := &textService.AddSubjectInClassRequest{
		Class: &textService.Class{
			Number:     req.Subject.ClassNumber,
			SubjectIds: []string{id.String()},
		},
	}

	err = postgresRepo.AddSubjectInClass(th.pg, updateRequest)
	if err != nil {
		return nil, fmt.Errorf("createSubject: %v", err)
	}

	return &textService.CreateSubjectResponse{
		Response: "Subject created successfully",
	}, nil
}

func (th *TextHandler) GetSubject(ctx context.Context, req *textService.GetSubjectRequest) (*textService.GetSubjectResponse, error) {
	// subject, err := redisRepo.GetSubject(ctx, th.redis, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("getSubject: %v", err)
	// }

	// if Subject == nil {
	subject, err := postgresRepo.SelectSubject(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("getSubject: %v", err)
	}

	return subject, nil
	// }

	// 	return Subject, nil
}

func (th *TextHandler) GetSubjects(ctx context.Context) (*textService.GetSubjectsResponse, error) {
	// Subjectes, err := redisRepo.GetSubjectes(th.redis, ctx)
	// if err != nil {
	// 	return nil, fmt.Errorf("getSubjectes: %v", err)
	// }

	// if len(Subjectes.Subjectes) == 0 {
	Subjectes, err := postgresRepo.SelectSubjects(th.pg)
	if err != nil {
		return nil, fmt.Errorf("getSubjectes: %v", err)
	}

	return Subjectes, nil
	// }

	// return Subjectes, nil
}

func (th *TextHandler) UpdateSubject(ctx context.Context, req *textService.UpdateSubjectRequest) (*textService.UpdateSubjectResponse, error) {
	// err := redisRepo.UpdateSubject(th.redis, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("updateSubject: %v", err)
	// }

	err := postgresRepo.UpdateSubject(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("updateSubject: %v", err)
	}

	return &textService.UpdateSubjectResponse{
		Response: "Subject updated successfully",
	}, nil
}
