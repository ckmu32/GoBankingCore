package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ckmu32/GoBankingCore/models"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

// AddAccount Let's us add a new account.
func AddAccount(account *models.Account) models.OperationResponse {

	//Check that the account does not exists.
	//Check that the cardHolder exists.

	/*
		account.Number = "125532634636"
		s, _ := json.Marshal(account)
		w.Header().Add("Content-Type", "application/json") //Must be before the http code assignation.
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(s))
	*/
	var errorOccurred = false

	response := models.OperationResponse{}

	if errorOccurred {
		response.Code = http.StatusUnprocessableEntity
		response.Description = "Errors ocurred."
		response.Errors = append(response.Errors, "Something weird happened.")
		return response
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

	response.Code = http.StatusCreated
	response.Description = account1
	/*
			{
		    "Code": 201,
		    "Description": "eyJudW1iZXIiOiIxMjQxMjM1NCIsImNhcmRfaG9sZGVyX2lkIjo3ODgyNiwidHlwZSI6IlBheWNoZWNrIiwiZXh0ZXJuYWxfYWNjb3VudCI6IjM0NjMzMjU2MzM0MzQ5OCIsImxldmVsIjoiNCIsImF2YWlsYWJsZV9iYWxhbmNlIjozNTAuNSwicmVhbF9iYWxhbmNlIjozMjAsInN0YXR1cyI6IkFDVElWRSIsImNyZWF0aW9uX2RhdGUiOiIiLCJjcmVhdGlvbl9ob3VyIjoiIiwiY3JlYXRpb25fdXNlciI6IiIsIm1vZmljYXRpb25fZGF0ZSI6IiIsIm1vZmljYXRpb25faG91ciI6IiIsIm1vZGlmaWNhdGlvbl91c2VyIjoiIn0=",
		    "Errors": null
		}
	*/

	/*


			{
		    "Code": 422,
		    "Description": "Errors ocurred.",
		    "Errors": [
		        "Something weird happened."
		    ]
		}
					{
				    "Code": 201,
				    "Description": {
							cuenta{
								"number": "12412354",
								"card_holder_id": 78826,
								"type": "Paycheck",
								"external_account": "346332563343498",
								"level": "4",
								"available_balance": 350.5,
								"real_balance": 320,
								"status": "ACTIVE",
								"creation_date": "",
								"creation_hour": "",
								"creation_user": "",
								"mofication_date": "",
								"mofication_hour": "",
								"modification_user": ""
							}
				    },
				    "Errors": null
				}

							{
				    "Code": 201,
				    "Description": {
				        "number": "12412354",
				        "card_holder_id": 78826,
				        "type": "Paycheck",
				        "external_account": "346332563343498",
				        "level": "4",
				        "available_balance": 350.5,
				        "real_balance": 320,
				        "status": "ACTIVE",
				        "creation_date": "",
				        "creation_hour": "",
				        "creation_user": "",
				        "mofication_date": "",
				        "mofication_hour": "",
				        "modification_user": ""
				    },
				    "Errors": null
				}
	*/

	return response
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
