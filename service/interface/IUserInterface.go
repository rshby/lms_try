package _interface

import (
	"context"
	"lms_try/model/dto"
)

type IUserService interface {
	GetAll(ctx context.Context) ([]dto.UserDetail, error)
	GetById(ctx context.Context, id *dto.UserIdRequest) ([]dto.UserDetail, error)
}
