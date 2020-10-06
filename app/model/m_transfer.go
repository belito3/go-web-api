package model

import "time"

// Simple bank example
type Transfer struct {
	ID				int64		`json:"id"`
	FromAccountID	int64		`json:"from_account_id"`
	ToAccountID		int64		`json:"to_account_id"`
	// must be positive
	Amount 			int64		`json:"amount"`
	CreatedAt		time.Time	`json:"created_at"`
}