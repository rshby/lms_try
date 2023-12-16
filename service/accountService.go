package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"lms_try/model/dto"
	_interface "lms_try/repository/interface"
	interfaceService "lms_try/service/interface"
)

type AccountService struct {
	AccountRepository _interface.IAccountRepository
	logger            *logrus.Entry
}

// function provider
func NewAccountService(accRepo _interface.IAccountRepository, log *logrus.Logger) interfaceService.IAccountService {
	return &AccountService{
		AccountRepository: accRepo,
		logger: log.WithFields(logrus.Fields{
			"Package": "service",
			"Class":   "AccountService",
		}),
	}
}

// mehtod Insert new data user
func (a AccountService) Insert(ctx context.Context, request *dto.InsertAccountRequest) error {
	//TODO implement me
	panic("implement me")
}
