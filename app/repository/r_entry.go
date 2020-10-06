package repository

import (
	"context"
	"github.com/belito3/go-api-codebase/app/model"
)

type CreateEntryParams struct {
	AccountID	int64	`json:"account_id"`
	Amount		int64	`json:"amount"`
}

type ListEntriesParams struct {
	AccountID 	int64	`json:"account_id"`
	Limit		int32	`json:"limit"`
	Offset		int32	`json:"offset"`
}

type Entry = model.Entry

type IEntry interface {
	CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
	GetEntry(ctx context.Context, id int64) (Entry, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
}