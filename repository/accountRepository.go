package repository

import (
	"context"
	"database/sql"
	lr "github.com/sirupsen/logrus"
	myError "lms_try/helper/error"
	"lms_try/model/entity"
	_interface "lms_try/repository/interface"
)

type AccountRepository struct {
	DB     *sql.DB
	logger *lr.Entry
}

// function provider
func NewAccountRepository(db *sql.DB, logrus *lr.Logger) _interface.IAccountRepository {
	return &AccountRepository{
		DB: db,
		logger: logrus.WithFields(lr.Fields{
			"Package": "Repository",
			"Class":   "AccountRepository",
		}),
	}
}

// method Insert data account to database
func (a *AccountRepository) Insert(ctx context.Context, entity *entity.Account) (*entity.Account, error) {
	a.logger.WithFields(lr.Fields{
		"method": "Insert",
	})

	// create prepare insert query
	query, err := a.DB.PrepareContext(ctx, "INSERT INTO accounts (email, password, user_id) VALUES (?, ?, ?)")
	if err != nil {
		a.logger.Error(err.Error())
		return nil, myError.NewServerError(err.Error())
	}

	// execute insert query
	result, err := query.ExecContext(ctx, entity.Email, entity.Password, entity.UserId)
	if err != nil {
		a.logger.Error(err.Error())
		return nil, myError.NewServerError(err.Error())
	}

	if row, _ := result.RowsAffected(); row == 0 {
		a.logger.Error("rows affected are zero")
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		a.logger.Error(err.Error())
		return nil, err
	}

	entity.Id = int(id)

	// success
	a.logger.Info("success insert")
	return entity, nil
}

// method get data account by email
func (a *AccountRepository) GetByEmail(ctx context.Context, email string) (*entity.Account, error) {
	//TODO implement me
	panic("implement me")
}
