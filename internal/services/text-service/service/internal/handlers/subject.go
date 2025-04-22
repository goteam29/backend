package handlers

import (
	postgresRepo "api-repository/internal/services/text-service/service/internal/repository/postgres"
	textService "api-repository/pkg/api/text-service"
	"context"
	"database/sql"
	"fmt"
)

func (th *TextHandler) CreateSubject(ctx context.Context, req *textService.CreateSubjectRequest) (*textService.CreateSubjectResponse, error) {
	// err := redisRepo.AddSubject(th.redis, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("createSubject: %v", err)
	// }

	tx, err := th.pg.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, fmt.Errorf("pgAddSubjectInClass: failed to begin transaction: %v", err)
	}
	defer tx.Rollback()

	id, err := postgresRepo.InsertSubject(ctx, th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("createSubject: %v", err)
	}

	updateRequest := &textService.AddSubjectInClassRequest{
		Id:        req.ClassId,
		SubjectId: id.Id,
	}

	_, err = postgresRepo.AddSubjectInClass(ctx, th.pg, updateRequest)
	if err != nil {
		return nil, fmt.Errorf("createSubject: %v", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("createSubject: failed to commit transaction: %v", err)
	}

	return id, nil
}

func (th *TextHandler) GetSubject(ctx context.Context, req *textService.GetSubjectRequest) (*textService.GetSubjectResponse, error) {
	// subject, err := redisRepo.GetSubject(ctx, th.redis, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("getSubject: %v", err)
	// }

	// if Subject == nil {
	subject, err := postgresRepo.SelectSubject(ctx, th.pg, req)
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
	subjects, err := postgresRepo.SelectSubjects(ctx, th.pg)
	if err != nil {
		return nil, fmt.Errorf("getSubjectes: %v", err)
	}

	return subjects, nil
	// }

	// return Subjectes, nil
}

func (th *TextHandler) AddSectionInSubject(ctx context.Context, req *textService.AddSectionInSubjectRequest) (*textService.AddSectionInSubjectResponse, error) {
	// err := redisRepo.UpdateSubject(th.redis, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("updateSubject: %v", err)
	// }

	tx := &sql.Tx{}

	sectionId, err := postgresRepo.AddSectionInSubject(ctx, tx, req)
	if err != nil {
		return nil, fmt.Errorf("updateSubject: %v", err)
	}

	return sectionId, nil
}

func (th *TextHandler) RemoveSectionFromSubject(ctx context.Context, req *textService.RemoveSectionFromSubjectRequest) (*textService.RemoveSectionFromSubjectResponse, error) {
	// err := redisRepo.UpdateSubject(th.redis, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("updateSubject: %v", err)
	// }

	sectionId, err := postgresRepo.RemoveSectionFromSubject(ctx, th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("updateSubject: %v", err)
	}

	return sectionId, nil
}

func (th *TextHandler) DeleteSubject(ctx context.Context, req *textService.DeleteSubjectRequest) (*textService.DeleteSubjectResponse, error) {
	// err := redisRepo.DeleteSubject(th.redis, req)
	// if err != nil {
	// 	return nil, fmt.Errorf("deleteSubject: %v", err)
	// }

	id, err := postgresRepo.DeleteSubject(ctx, th.pg, req)
	if err != nil {
		return nil, fmt.Errorf("deleteSubject: %v", err)
	}

	return id, nil
}
