package handlers

import (
	postgresRepo "api-repository/internal/services/text-service/service/internal/repository/postgres"
	textService "api-repository/pkg/api/text-service"
	"context"
	"fmt"
)

func (th *TextHandler) CreateSection(ctx context.Context, req *textService.CreateSectionRequest) (*textService.CreateSectionResponse, error) {
	id, err := postgresRepo.InsertSection(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("insertSection: %v", err)
	}

	return id, nil
}

func (th *TextHandler) GetSection(ctx context.Context, req *textService.GetSectionRequest) (*textService.GetSectionResponse, error) {
	section, err := postgresRepo.SelectSection(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("getSection: %v", err)
	}

	return section, nil
}

func (th *TextHandler) GetSections(ctx context.Context) (*textService.GetSectionsResponse, error) {
	sections, err := postgresRepo.SelectSections(th.pg)
	if err != nil {
		return nil, fmt.Errorf("getSections: %v", err)
	}

	return sections, nil
}

func (th *TextHandler) AssignLessonToSection(ctx context.Context, req *textService.AssignLessonToSectionRequest) (*textService.AssignLessonToSectionResponse, error) {
	lessonId, err := postgresRepo.UpdateSection(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("addLessonsInSection: %v", err)
	}

	return lessonId, nil
}

func (th *TextHandler) DeleteSection(ctx context.Context, req *textService.DeleteSectionRequest) (*textService.DeleteSectionResponse, error) {
	id, err := postgresRepo.DeleteSection(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("deleteSection: %v", err)
	}

	return id, nil
}
