package repository

import (
	"context"
	"github.com/belito3/go-api-codebase/app/model"
)

type Account = model.Account

type AddAccountBalanceParams struct {
	Amount int64 `json:"amount"`
	ID     int64 `json:"id"`
}

type CreateAccountParams struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

type ListAccountsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type UpdateAccountParams struct {
	ID      int64 `json:"id"`
	Balance int64 `json:"balance"`
}

type IAccount interface {
	AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error)
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	GetAccount(ctx context.Context, id int64) (Account, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	DeleteAccount(ctx context.Context, id int64) error
	GetAccountForUpdate(ctx context.Context, id int64) (Account, error)
	ListAccount(ctx context.Context, arg ListAccountsParams) ([]Account, error)
}
