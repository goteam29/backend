package handlers

import (
	postgresRepo "api-repository/internal/services/text-service/service/internal/repository/postgres"
	textService "api-repository/pkg/api/text-service"
	"context"
	"fmt"

	uuid "github.com/google/uuid"
)

func (th *TextHandler) CreateSection(ctx context.Context, req *textService.CreateSectionRequest) (*textService.CreateSectionResponse, error) {
	// err := redisRepo.InsertSection(th.redis, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("insertSection: %v", err)
	// }

	id := uuid.New()

	err := postgresRepo.InsertSection(th.pg, req, id)
	if err != nil {
		return nil, fmt.Errorf("insertSection: %v", err)
	}

	return &textService.CreateSectionResponse{
		Response: "section created successfully",
	}, nil
}

// func (th *TextHandler) GetSection(ctx context.Context, req *textService.GetSectionRequest) (*textService.GetSectionResponse, error) {
// 	// section, err := redisRepo.GetSection(th.redis, req)
// 	// if err != nil {
// 	// 	return nil, fmt.Errorf("getSection: %v", err)
// 	// }

// 	section, err := postgresRepo.SelectSection(th.pg, req)
// 	if err != nil {
// 		return nil, fmt.Errorf("getSection: %v", err)
// 	}

// 	return section, nil
// }

// func (th *TextHandler) GetSections(ctx context.Context) (*textService.GetSectionsResponse, error) {
// 	// sections, err := redisRepo.GetSections(th.redis, ctx)
// 	// if err != nil {
// 	// 	return nil, fmt.Errorf("getSections: %v", err)
// 	// }

// 	sections, err := postgresRepo.SelectSections(th.pg)
// 	if err != nil {
// 		return nil, fmt.Errorf("getSections: %v", err)
// 	}

// 	return sections, nil
// }

// func (th *TextHandler) AddLessonsInSection(ctx context.Context, req *textService.AddLessonsInSectionRequest) (*textService.AddLessonsInSectionResponse, error) {
// 	// err := redisRepo.AddLessonsInSection(th.redis, req)
// 	// if err != nil {
// 	// 	return nil, fmt.Errorf("addLessonsInSection: %v", err)
// 	// }

// 	err := postgresRepo.AddLessonsInSection(th.pg, req)
// 	if err != nil {
// 		return nil, fmt.Errorf("addLessonsInSection: %v", err)
// 	}

// 	return &textService.AddLessonsInSectionResponse{
// 		Response: "lessons added in section successfully",
// 	}, nil
// }

// func (th *TextHandler) RemoveLessonsFromSection(ctx context.Context, req *textService.RemoveLessonFromSectionRequest) (*textService.RemoveLessonFromSectionResponse, error) {
// 	// err := redisRepo.RemoveLessonsFromSection(th.redis, req)
// 	// if err != nil {
// 	// 	return nil, fmt.Errorf("removeLessonsFromSection: %v", err)
// 	// }

// 	err := postgresRepo.RemoveLessonFromSection(th.pg, req)
// 	if err != nil {
// 		return nil, fmt.Errorf("removeLessonFromSection: %v", err)
// 	}

// 	return &textService.RemoveLessonFromSectionResponse{
// 		Response: "lesson removed from section successfully",
// 	}, nil
// }

// func (th *TextHandler) DeleteSection(ctx context.Context, req *textService.DeleteSectionRequest) (*textService.DeleteSectionResponse, error) {
// 	// err := redisRepo.DeleteSection(th.redis, req)
// 	// if err != nil {
// 	// 	return nil, fmt.Errorf("deleteSection: %v", err)
// 	// }

// 	err := postgresRepo.DeleteSection(th.pg, req)
// 	if err != nil {
// 		return nil, fmt.Errorf("deleteSection: %v", err)
// 	}

// 	return &textService.DeleteSectionResponse{
// 		Response: "section deleted successfully",
// 	}, nil
// }
