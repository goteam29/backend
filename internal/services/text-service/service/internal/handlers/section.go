package handlers

import (
	postgresRepo "api-repository/internal/services/text-service/service/internal/repository/postgres"
	textService "api-repository/pkg/api/text-service"
	"context"
	"database/sql"
	"fmt"
)

func (th *TextHandler) CreateSection(ctx context.Context, req *textService.CreateSectionRequest) (*textService.CreateSectionResponse, error) {
	// err := redisRepo.InsertSection(th.redis, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("insertSection: %v", err)
	// }

	tx, err := th.pg.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("pgAddSubjectInClass: failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	id, err := postgresRepo.InsertSection(tx, req)
	if err != nil {
		return nil, fmt.Errorf("insertSection: %v", err)
	}

	updateRequest := &textService.AddSectionInSubjectRequest{
		Id:        req.SubjectId,
		SectionId: id.Id,
	}

	_, err = postgresRepo.AddSectionInSubject(ctx, tx, updateRequest)
	if err != nil {
		return nil, fmt.Errorf("createSection: failed to add section in subject: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("createSection: failed to commit transaction: %v", err)
	}

	return id, nil
}

func (th *TextHandler) GetSection(ctx context.Context, req *textService.GetSectionRequest) (*textService.GetSectionResponse, error) {
	// section, err := redisRepo.GetSection(th.redis, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("getSection: %v", err)
	// }

	section, err := postgresRepo.SelectSection(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("getSection: %v", err)
	}

	return section, nil
}

func (th *TextHandler) GetSections(ctx context.Context) (*textService.GetSectionsResponse, error) {
	// sections, err := redisRepo.GetSections(th.redis, ctx)
	// if err != nil {
	// 	return nil, fmt.Errorf("getSections: %v", err)
	// }

	sections, err := postgresRepo.SelectSections(th.pg)
	if err != nil {
		return nil, fmt.Errorf("getSections: %v", err)
	}

	return sections, nil
}

func (th *TextHandler) AddLessonInSection(ctx context.Context, req *textService.AddLessonInSectionRequest) (*textService.AddLessonInSectionResponse, error) {
	// err := redisRepo.AddLessonsInSection(th.redis, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("addLessonsInSection: %v", err)
	// }

	lessonId, err := postgresRepo.AddLessonInSection(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("addLessonsInSection: %v", err)
	}

	return lessonId, nil
}

func (th *TextHandler) RemoveLessonFromSection(ctx context.Context, req *textService.RemoveLessonFromSectionRequest) (*textService.RemoveLessonFromSectionResponse, error) {
	// err := redisRepo.RemoveLessonsFromSection(th.redis, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("removeLessonsFromSection: %v", err)
	// }

	lessonId, err := postgresRepo.RemoveLessonFromSection(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("removeLessonFromSection: %v", err)
	}

	return lessonId, nil
}

func (th *TextHandler) DeleteSection(ctx context.Context, req *textService.DeleteSectionRequest) (*textService.DeleteSectionResponse, error) {
	// err := redisRepo.DeleteSection(th.redis, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("deleteSection: %v", err)
	// }

	id, err := postgresRepo.DeleteSection(th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("deleteSection: %v", err)
	}

	return id, nil
}
