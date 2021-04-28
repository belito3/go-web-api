package repository

import (
	"context"
)

// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult is the result of the transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// Store define all functions to execute db queries and transaction
type IStore interface {
	IAccount
	IEntry
	ITransfer
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
	TransferTx2(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}