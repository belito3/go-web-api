package impl

import (
	"context"
	"github.com/belito3/go-api-codebase/app/model"
	"github.com/belito3/go-api-codebase/pkg/errors"
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


func (a *Queries) AddAccountBalance(ctx context.Context, arg AddAccountBalanceParams) (Account, error) {
	query := `UPDATE accounts
		SET balance = balance + $1
		WHERE id = $2
		RETURNING id, owner, balance, currency, created_at`

	row := a.db.QueryRowContext(ctx, query, arg.Amount, arg.ID)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

// TODO: insert return value ở oracle phải dùng script procedure lằng nhằng quá
// nên chuyển qua dbsql
func (a *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	query := `INSERT INTO accounts (owner, balance, currency) 
			VALUES ($1, $2, $3) RETURNING *`

	var i Account

	row := a.db.QueryRowContext(ctx, query, arg.Owner, arg.Balance, arg.Currency)
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

func (a *Queries) GetAccount(ctx context.Context, id int64) (Account, error) {
	query := `SELECT id, owner, balance, currency, created_at FROM accounts 
			WHERE id = $1 LIMIT 1`
	var i Account
	row := a.db.QueryRowContext(ctx, query, id)
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

func (a *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	query := `UPDATE accounts SET balance = $2 WHERE id = $1 
			RETURNING id, owner, balance, currency, created_at`

	row := a.db.QueryRowContext(ctx, query, arg.ID, arg.Balance)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

func (a *Queries) DeleteAccount(ctx context.Context, id int64) error {
	query := `DELETE FROM accounts WHERE id = $1`
	_, err := a.db.ExecContext(ctx, query, id)
	return err
}

func (a *Queries) GetAccountForUpdate(ctx context.Context, id int64) (Account, error) {
	/*
		https://www.cockroachlabs.com/docs/stable/select-for-update.html#required-privileges
		SELECT FOR UPDATE: Enforce transaction order when updating the same rows
		In this example, we'll use SELECT FOR UPDATE to lock a row inside a transaction, forcing other
		transactions that want update same row to wait for the first transaction to complete. The other
		transactions that want to update the same row are effectively put into a queue based on when they
		first try to read the value of the row
	*/
	// This tell postgres that we don't update the key, or column of accounts table
	query := `SELECT id, owner, balance, currency, created_at FROM accounts 
			WHERE id = $1 LIMIT 1 FOR NO KEY UPDATE`
	row := a.db.QueryRowContext(ctx, query, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

func (a *Queries) ListAccount(ctx context.Context, arg ListAccountsParams) ([]Account, error) {
	query := `SELECT id, owner, balance, currency, created_at FROM accounts 
			ORDER BY id LIMIT $1 OFFSET $2`
	// query := "SELECT id, owner, balance, currency, created_at FROM accounts ORDER BY id LIMIT $1 OFFSET $2"
	rows, err := a.db.QueryContext(ctx, query, arg.Limit, arg.Offset)
	if err != nil {
		return nil, errors.Wrapf(err, "QueryContext ")
		// return nil, err
	}
	defer rows.Close()
	var items []Account
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Balance,
			&i.Currency,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
