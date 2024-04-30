package main

import (
	"bank/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/accounts/all", handlers.GetAccounts)
	http.HandleFunc("/accounts/{account_number}", handlers.GetAccountByAccountNumber)
	http.HandleFunc("/accounts/custmorid/{customer_id}", handlers.GetAccountByCustomerID)
	http.HandleFunc("/accounts/new", handlers.CreateAccount)
	http.HandleFunc("/accounts/update/{account_number}", handlers.UpdateAccount)
	http.HandleFunc("/accounts/delete/{account_number}", handlers.DeleteAccount)

	// Customer Apies
	http.HandleFunc("/customer/new", handlers.CreatCustomer)
	http.HandleFunc("/customer/all", handlers.AllCustomer)
	http.HandleFunc("/customer/search/{id}", handlers.SearchById)
	http.HandleFunc("/customer/deletebyid/{id}", handlers.DeleteById)
	http.HandleFunc("/customer/updatebyid/{id}", handlers.UpdateById)

	// Test Apies
	http.HandleFunc("/allpost", handlers.GetAllPost)
	http.HandleFunc("/searchbyid/{id}", handlers.GetPostByid)
	http.HandleFunc("/updatebyid/{id}", handlers.UpdatePostById )
	http.ListenAndServe(":8080", nil)

}
 