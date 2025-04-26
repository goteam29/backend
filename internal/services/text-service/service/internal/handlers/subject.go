package handlers

import (
	postgresRepo "api-repository/internal/services/text-service/service/internal/repository/postgres"
	textService "api-repository/pkg/api/text-service"
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (th *TextHandler) CreateSubject(ctx context.Context, req *textService.CreateSubjectRequest) (*textService.CreateSubjectResponse, error) {
	id, err := postgresRepo.InsertSubject(ctx, th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("createSubject: %v", err)
	}

	return id, nil
}

func (th *TextHandler) GetSubject(ctx context.Context, req *textService.GetSubjectRequest) (*textService.GetSubjectResponse, error) {
	subject, err := postgresRepo.SelectSubject(ctx, th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("getSubject: %v", err)
	}

	return subject, nil
}

func (th *TextHandler) GetSubjects(ctx context.Context) (*textService.GetSubjectsResponse, error) {
	subjects, err := postgresRepo.SelectSubjects(ctx, th.pg)
	if err != nil {
		return nil, fmt.Errorf("getSubjectes: %v", err)
	}

	return subjects, nil
}

func (th *TextHandler) AssignSectionToSubject(ctx context.Context, req *textService.AssignSectionToSubjectRequest) (*textService.AssignSectionToSubjectResponse, error) {
	sectionId, err := postgresRepo.UpdateSubject(ctx, th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("updateSubject: %v", err)
	}

	return sectionId, nil
}

func (th *TextHandler) DeleteSubject(ctx context.Context, req *textService.DeleteSubjectRequest) (*emptypb.Empty, error) {
	id, err := postgresRepo.DeleteSubject(ctx, th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("deleteSubject: %v", err)
	}

	return id, nil
}
