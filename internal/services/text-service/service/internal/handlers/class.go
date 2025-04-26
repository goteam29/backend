package handlers

import (
	postgresRepo "api-repository/internal/services/text-service/service/internal/repository/postgres"
	textService "api-repository/pkg/api/text-service"
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (th *TextHandler) CreateClass(ctx context.Context, req *textService.CreateClassRequest) (*textService.CreateClassResponse, error) {
	id, err := postgresRepo.InsertClass(ctx, th.pg, req)
	if err != nil {
		if errors.Is(err, postgresRepo.ErrClassAlreadyExists) {
			return nil, status.Errorf(codes.AlreadyExists, "class with number %d already exists", req.Number)
		}
		return nil, status.Errorf(codes.Internal, "createClass: failed: %v", err)
	}

	return id, nil
}

func (th *TextHandler) GetClass(ctx context.Context, req *textService.GetClassRequest) (*textService.GetClassResponse, error) {
	class, err := postgresRepo.SelectClass(ctx, th.pg, req)
	if err != nil {
		if errors.Is(err, postgresRepo.ErrClassNotFound) {
			return nil, status.Errorf(codes.NotFound, "class with id %s not found", req.Id)
		}
		return nil, status.Errorf(codes.Internal, "getClass: failed to get class: %v", err)
	}

	return class, nil
}

func (th *TextHandler) GetClasses(ctx context.Context) (*textService.GetClassesResponse, error) {
	classes, err := postgresRepo.SelectClasses(ctx, th.pg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "getClasses: failed to get classes: %v", err)
	}

	return classes, nil
}

func (th *TextHandler) AddSubjectInClass(ctx context.Context, req *textService.AddSubjectInClassRequest) (*textService.AddSubjectInClassResponse, error) {
	subjectId, err := postgresRepo.AddSubjectInClass(ctx, th.pg, req)
	if err != nil {
		if errors.Is(err, postgresRepo.ErrClassOrSubjectNotFound) {
			return nil, status.Errorf(codes.NotFound, "class or subject not found")
		}
		return nil, status.Errorf(codes.Internal, "addSubjectInClass: failed: %v", err)
	}

	return subjectId, nil
}

func (th *TextHandler) RemoveSubjectFromClass(ctx context.Context, req *textService.RemoveSubjectFromClassRequest) (*textService.RemoveSubjectFromClassResponse, error) {
	subjectId, err := postgresRepo.RemoveSubjectFromClass(ctx, th.pg, req)
	if err != nil {
		if errors.Is(err, postgresRepo.ErrSubjectNotFound) {
			return nil, status.Errorf(codes.NotFound, "class or subject not found")
		}
		return nil, status.Errorf(codes.Internal, "removeSubjectFromClass: failed: %v", err)
	}

	return subjectId, nil
}

func (th *TextHandler) DeleteClass(ctx context.Context, req *textService.DeleteClassRequest) (*emptypb.Empty, error) {
	id, err := postgresRepo.DeleteClass(ctx, th.pg, req)
	if err != nil {
		if errors.Is(err, postgresRepo.ErrClassNotFound) {
			return nil, status.Errorf(codes.NotFound, "class not found")
		}
		return nil, status.Errorf(codes.Internal, "deleteClass: failed: %v", err)
	}

	return id, nil
}
