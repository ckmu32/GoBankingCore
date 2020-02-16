package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ckmu32/GoBankingCore/models"
	"github.com/ckmu32/GoBankingCore/repositories"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// AddAccount Let's us add a new account.
func AddAccount(account *models.Account) models.OperationResponse {

	//Check that the account does not exists.
	//Check that the cardHolder exists.

	response := models.OperationResponse{}

	if exists(account.Number, &response) {
		return response
	}

	/* TRASH
	var errorOccurred = false
	if errorOccurred {
		response.Code = http.StatusUnprocessableEntity
		response.Description = "Errors ocurred."
		response.Errors = append(response.Errors, "Something weird happened.")
		return response
	}
	*/

	//TEMP
	account.CreationUser = "TELLER_0"
	account.ModificationUser = "TELLER_O"

	// account1 := models.Account{
	// 	Level:            "4",
	// 	Type:             "Paycheck",
	// 	CardHolderID:     78826,
	// 	ExternalAccount:  "346332563343498",
	// 	Number:           "12412354",
	// 	Status:           "ACTIVE",
	// 	AvailableBalance: 350.50,
	// 	RealBalance:      320,
	// }

	var inserted = repositories.Insert(account, &response)

	if inserted {
		response.Code = http.StatusCreated
		response.Description = account
	}

	return response
}

// exists Checks if the account exists on the DB.
func exists(number string, response *models.OperationResponse) bool {
	var retrieved = repositories.GetByNumber(number, response)
	// This is the only code we used to check if everything is correct.
	if response.Code == http.StatusNotFound {
		response.Errors = nil
		return false
	}

	// If the record was retrieved, we check by comparing the account and if its equal we set a specific error.
	if retrieved.Number == number {
		response.Code = http.StatusUnprocessableEntity
		response.Description = "Errors occurred."
		response.Errors = append(response.Errors, "The account already exists.")
	}
	return true
}

// GetAccount Obtain one account according to the account number. Responds to GET.
func GetAccount(number string, response *models.OperationResponse) models.Account {

	retrievedAccount := repositories.GetByNumber("1212121212", response)

	return retrievedAccount
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

// UpdateAccount Update the account using the number and the attributes of the object. For the purpose of history. Responds to POST.
func UpdateAccount(w http.ResponseWriter, r *http.Request) {
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

	//This with the declaration up ahead
	validate = validator.New()
	//Or this way witout declaration.
	//validate _= validator.New()
	structError := validate.Struct(account)

	for _, e := range structError.(validator.ValidationErrors) {
		fmt.Println(e)
	}

	w.Write([]byte("Record updated"))
}

func validAccount(account *models.Account, response *models.OperationResponse) bool {
	validate = validator.New()
	//Or this way witout declaration.
	//validate _= validator.New()
	structError := validate.Struct(account)

	for _, e := range structError.(validator.ValidationErrors) {
		fmt.Println(e)
	}

	return true
}
