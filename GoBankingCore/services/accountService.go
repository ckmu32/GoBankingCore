package services

import (
	"encoding/json"
	"net/http"

	"github.com/ckmu32/GoBankingCore/models"
)

// AddAccount Let's us add a new account. Responds to POST.
func AddAccount(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	account := models.Account{}
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Unprocessable Entity"))
		return
	}
	account.Number = "125532634636"
	s, _ := json.Marshal(account)
	w.Header().Add("Content-Type", "application/json") //Must be before the http code assignation.
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(s))
}

// GetAccount Obtain one account according to the account number. Responds to GET.
func GetAccount(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	account1 := models.Account{
		Level:            "4",
		Type:             "Paycheck",
		CardHolderID:     78826,
		ExternalAccount:  "346332563343498",
		Number:           "12412354",
		Status:           "ACTIVE",
		AvailableBalance: 350.50,
		RealBalance:      320,
	}

	s, _ := json.Marshal(account1)
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(s))
}

// GetAccounts Obtain all accounts . Responds to GET.
func GetAccounts(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	var accountArrayString []string
	accountArrayString = append(accountArrayString, "Cuenta 1")
	accountArrayString = append(accountArrayString, "Cuenta 2")

	var accountArray []models.Account
	account1 := models.Account{
		Level:            "4",
		Type:             "Paycheck",
		CardHolderID:     78826,
		ExternalAccount:  "346332563343498",
		Number:           "12412354",
		Status:           "ACTIVE",
		AvailableBalance: 350.50,
		RealBalance:      320,
	}
	account2 := models.Account{
		Level:            "3",
		Type:             "Investment",
		CardHolderID:     78826,
		ExternalAccount:  "346338243343411",
		Number:           "00412300",
		Status:           "ACTIVE",
		AvailableBalance: 1000,
		RealBalance:      1000,
	}
	accountArray = append(accountArray, account1)
	accountArray = append(accountArray, account2)

	s, _ := json.Marshal(accountArray)
	w.Header().Add("Content-Type", "application/json")
	w.Write([]byte(s))
}

// CancelAccount Cancels and account (Does not delete the record). For the purpose of history. Responds to POST.
func CancelAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
		return
	}

	account := models.Account{}
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte("Unprocessable Entity"))
		return
	}

	account.Status = "CANCELED"
	w.Write([]byte("Record canceled"))
}
