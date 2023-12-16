package _interface

import (
	"context"
	"lms_try/model/dto"
)

type IAccountService interface {
	Insert(ctx context.Context, request *dto.InsertAccountRequest) error
	// TODO : create service accountService implement this interface
}
