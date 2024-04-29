package models

type Customer struct {
	ID            int    `json:"id"`
	AccountNumber int    `json:"account_number"`
	Name          string `json:"name"`
	DOB           string `json:"dob"`
	Address       string `json:"address"`
}
