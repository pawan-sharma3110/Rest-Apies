package handlers

import (
	"bank/models"
	"encoding/json"
	"net/http"
	"strconv"
)

var (
	accounts []models.Account
	// banks        []Bank
	// customers    []Customer
	// transactions []Transaction
)

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accounts)
}

func GetAccountByAccountNumber(w http.ResponseWriter, r *http.Request) {

	// aa := r.URL.Query().Get("account_number")

	aa := r.PathValue("account_number")

	println("account number is : ", aa)
	accountNumber, _ := strconv.Atoi(aa)

	for _, account := range accounts {
		if account.AccountNumber == accountNumber {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusFound)
			json.NewEncoder(w).Encode(account)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": "Account not found"})

}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account models.Account
	json.NewDecoder(r.Body).Decode(&account)
	accounts = append(accounts, account)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(account)

}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	// aa := r.URL.Query().Get("account_number")

	aa := r.PathValue("account_number")

	println("account number is : ", aa)
	accountNumber, _ := strconv.Atoi(aa)

	for i, account := range accounts {
		if account.AccountNumber == accountNumber {

			var data models.Account
			json.NewDecoder(r.Body).Decode(&data)
			
			accounts[i].Balance = data.Balance
			accounts[i].AccountType = data.AccountType
			accounts[i].UpdatedAt = data.UpdatedAt
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusFound)
			json.NewEncoder(w).Encode(account)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": "Account not found"})
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	aa := r.PathValue("account_number")

	println("account number is : ", aa)
	accountNumber, _ := strconv.Atoi(aa)

	for i, account := range accounts {
		if account.AccountNumber == accountNumber {

			var data models.Account
			json.NewDecoder(r.Body).Decode(&data)
			accounts= append(accounts[:i], accounts[i+1:]...)
			
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(accounts)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"message": "Account not found"})
}

