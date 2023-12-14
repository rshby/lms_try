package _interface

import (
	"context"
	"lms_try/model/entity"
	"sync"
)

type IUserRepository interface {
	GetAll(ctx context.Context) ([]entity.User, error)
	GetByIdAsync(ctx context.Context, wg sync.WaitGroup, id int, user chan entity.User)
}
