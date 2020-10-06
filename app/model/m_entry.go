package model

import "time"

// Simple bank example
type Entry struct {
	ID			int64		`json:"id"`
	AccountID	int64		`json:"account_id"`
	// can be negative or positive
	Amount		int64		`json:"amount"`
	CreatedAt	time.Time	`json:"created_at"`
}