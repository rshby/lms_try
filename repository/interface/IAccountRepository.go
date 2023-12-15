package _interface

import (
	"context"
	"lms_try/model/entity"
)

type IAccountRepository interface {
	Insert(ctx context.Context, entity *entity.Account) (*entity.Account, error)
	GetByEmail(ctx context.Context, email string) (*entity.Account, error)
}
