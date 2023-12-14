package service

import (
	"context"
	"github.com/sirupsen/logrus"
	myError "lms_try/helper/error"
	"lms_try/model/dto"
	"lms_try/model/entity"
	interfaceRepo "lms_try/repository/interface"
	interfaceService "lms_try/service/interface"
	"reflect"
	"sync"
)

type UserService struct {
	UserRepo interfaceRepo.IUserRepository
}

// function provider
func NewUserService(userRepo interfaceRepo.IUserRepository) interfaceService.IUserService {
	return &UserService{userRepo}
}

// method get all users
func (u *UserService) GetAll(ctx context.Context) ([]dto.UserDetail, error) {
	// call procedure from repo
	users, err := u.UserRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var results []dto.UserDetail
	for _, user := range users {
		result := dto.UserDetail{
			Id: user.Id,
		}

		if user.FirstName.Valid {
			result.FirstName = user.FirstName.String
		}

		if user.LastName.Valid {
			result.LastName = user.LastName.String
		}

		if user.Gender.Valid {
			result.Gender = user.Gender.String
		}

		if user.BirthDate.Valid {
			result.BirthDate = user.BirthDate.Time.Format("2006-01-02")
		}

		if user.AddressId.Valid {
			result.AddressId = int(user.AddressId.Int64)
		}

		if user.EducationId.Valid {
			result.EducationId = int(user.EducationId.Int64)
		}

		results = append(results, result)
	}

	return results, nil
}

// method get user by id async
func (u *UserService) GetById(ctx context.Context, id *dto.UserIdRequest) ([]dto.UserDetail, error) {
	logrus.Info("method service GetById")

	var wg sync.WaitGroup

	if len(id.Id) == 0 {
		return nil, myError.NewBadRequestError("request cant be empty")
	}

	users := make(chan entity.User, len(id.Id))

	// call method in repository with async
	for _, id := range id.Id {
		go u.UserRepo.GetByIdAsync(ctx, wg, id, users)
	}

	wg.Wait()

	i := 1
	var results []dto.UserDetail
	for user := range users {
		if reflect.DeepEqual(user, entity.User{}) {
			if i == len(id.Id) {
				close(users)
			}
			i++
			continue
		}

		var result dto.UserDetail
		result.Id = user.Id
		if user.FirstName.Valid {
			result.FirstName = user.FirstName.String
		}

		results = append(results, result)

		if i == len(id.Id) {
			close(users)
		}
		i++
	}

	// check if not found
	if len(results) == 0 {
		logrus.Error("record not found")
		return nil, myError.NewNotFoundError("record not found")
	}

	// success get data
	return results, nil
}
