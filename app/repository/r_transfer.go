package repository

import (
	"context"
	"github.com/belito3/go-api-codebase/app/model"
)

type Transfer = model.Transfer

type CreateTransferParams struct {
	FromAccountID 	int64	`json:"from_account_id"`
	ToAccountID		int64	`json:"to_account_id"`
	Amount			int64	`json:"amount"`
}

type ListTransfersParams struct {
	FromAccountID 	int64	`json:"from_account_id"`
	ToAccountID		int64	`json:"to_account_id"`
	Limit			int32	`json:"limit"`
	Offset			int32	`json:"offset"`
}

type ITransfer interface {
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	GetTransfer(ctx context.Context, id int64) (Transfer, error)
	ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error)
}