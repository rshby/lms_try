package repository

import (
	"context"
	"database/sql"
	"github.com/sirupsen/logrus"
	myError "lms_try/helper/error"
	"lms_try/model/entity"
	_interface "lms_try/repository/interface"
	"sync"
)

type UserRepository struct {
	DB *sql.DB
}

// create function provider
func NewUserRepository(db *sql.DB) _interface.IUserRepository {
	return &UserRepository{DB: db}
}

// method GetAll data users
func (u UserRepository) GetAll(ctx context.Context) ([]entity.User, error) {
	query, err := u.DB.PrepareContext(ctx, "SELECT id, first_name, last_name, gender, birthdate, address_id, education_id FROM users")
	if err != nil {
		return nil, err
	}

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, myError.NewServerError(err.Error())
	}

	var users []entity.User
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Gender, &user.BirthDate, &user.AddressId, &user.EducationId); err != nil {
			return nil, myError.NewServerError(err.Error())
		}

		users = append(users, user)
	}

	// if not found
	if len(users) == 0 {
		return nil, myError.NewNotFoundError("record not found")
	}

	return users, nil
}

// method get data user by id
func (u *UserRepository) GetByIdAsync(ctx context.Context, wg sync.WaitGroup, id int, user chan entity.User) {
	wg.Add(1)
	defer wg.Done()

	query, err := u.DB.PrepareContext(ctx, "SELECT id, first_name, last_name, gender, birthdate, address_id, education_id FROM users WHERE id=?")
	if err != nil {
		logrus.Error("error cant create prepare query :", err.Error())
		user <- entity.User{}
		return
	}

	row := query.QueryRowContext(ctx, id)
	if row.Err() != nil {
		logrus.WithField("method", "repository GetByIdAsync").Error("error cant exec query :", err.Error())
		user <- entity.User{}
		return
	}

	var result entity.User
	if err := row.Scan(&result.Id, &result.FirstName, &result.LastName, &result.Gender, &result.BirthDate, &result.AddressId, &result.EducationId); err != nil {
		logrus.WithField("method", "repository GetByIdAsync").Error("error cant scan :", err.Error())
		user <- entity.User{}
		return
	}

	// success get data
	logrus.WithField("method", "repository GetByIdAsync").Info("success get data with id ", id)
	user <- result
}
