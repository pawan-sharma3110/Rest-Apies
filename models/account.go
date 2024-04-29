package models

import "time"

type Account struct {
	AccountNumber int    `json:"account_number"`
	CustomerID    int    `json:"customer_id"`
	Balance       int    `json:"balance"`
	AccountType   string `json:"account_type"`
	// Transactions  []Transaction `json:"transactions"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
