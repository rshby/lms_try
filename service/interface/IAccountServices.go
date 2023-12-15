package _interface

import "context"

type IAccountService interface {
	Insert(ctx context.Context) error
	// TODO : create service accountService implement this interface
}
