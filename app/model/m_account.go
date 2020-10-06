package model

import "time"

// Simple bank example
// https://github.com/techschool/simplebank/blob/master/db/sqlc/models.go
type Account struct {
	ID 			int64		`json:"id"`
	Owner		string		`json:"owner"`
	Balance 	int64		`json:"balance"`
	Currency	string		`json:"currency"`
	CreatedAt	time.Time	`json:"created_at"`
}
