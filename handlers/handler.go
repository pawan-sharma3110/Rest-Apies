package handlers

import (
	"bank/models"
	"encoding/json"
	"net/http"
	"strconv"
)

// var (

// 	// banks        []Bank
// 	// customers    []Customer
// 	// transactions []Transaction
// )

var accounts []models.Account = []models.Account{
	{
		AccountNumber: 1001,
		CustomerID:    1,
		Balance:       5000,
		AccountType:   "Savings",
	},
	{
		AccountNumber: 1002,
		CustomerID:    2,
		Balance:       7000,
		AccountType:   "Checking",
	},
	{
		AccountNumber: 1003,
		CustomerID:    3,
		Balance:       3000,
		AccountType:   "Savings",
	},
	{
		AccountNumber: 1004,
		CustomerID:    4,
		Balance:       10000,
		AccountType:   "Checking",
	},
	{
		AccountNumber: 1005,
		CustomerID:    5,
		Balance:       2500,
		AccountType:   "Savings",
	},
	{
		AccountNumber: 1006,
		CustomerID:    6,
		Balance:       8000,
		AccountType:   "Checking",
	},
	{
		AccountNumber: 1007,
		CustomerID:    7,
		Balance:       4000,
		AccountType:   "Savings",
	},
	{
		AccountNumber: 1008,
		CustomerID:    8,
		Balance:       6000,
		AccountType:   "Checking",
	},
	{
		AccountNumber: 1009,
		CustomerID:    9,
		Balance:       1500,
		AccountType:   "Savings",
	},
	{
		AccountNumber: 1010,
		CustomerID:    10,
		Balance:       9000,
		AccountType:   "Checking",
	},
}

func GetAccounts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(accounts)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "You can only get reqest"})
	}
}

func GetAccountByAccountNumber(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
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

	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Account not found"})

	}

}
func GetAccountByCustomerID(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		aa := r.PathValue("customer_id")

		println("account number is : ", aa)
		peramsId, _ := strconv.Atoi(aa)

		for _, account := range accounts {
			if account.CustomerID == peramsId {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(account)
				return
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(map[string]string{"message": "Account not found."})
			}
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Account not found,Please make an get request."})
	}

}
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var account models.Account
		json.NewDecoder(r.Body).Decode(&account)
		accounts = append(accounts, account)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(account)

	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Account not created"})

	}
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {
	// aa := r.URL.Query().Get("account_number")

	if r.Method == http.MethodPatch {
		aa := r.PathValue("account_number")
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
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "Account not found so you cant update account"})
	}

}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {
		aa := r.PathValue("account_number")
		accountNumber, _ := strconv.Atoi(aa)

		for i, account := range accounts {
			if account.AccountNumber == accountNumber {

				var data models.Account
				json.NewDecoder(r.Body).Decode(&data)
				accounts = append(accounts[:i], accounts[i+1:]...)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(accounts)
				return
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(map[string]string{"message": "Account not found so you cant delete account"})
			}
		}
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": "You can only delete reqeste"})
	}

}
