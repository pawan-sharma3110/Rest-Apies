package models

type Transaction struct {
	AccountNumber   int    `json:"account_number"`
	CustomerID      int    `json:"customer_id"`
	Status          string `json:"status"`
	Amount          int    `json:"amount"`
	TransactionType string `json:"transaction_type"`
	FromAccount     string `json:"from"`
	ToAccount       string `json:"to"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}
