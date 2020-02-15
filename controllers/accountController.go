package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ckmu32/GoBankingCore/models"
	"github.com/ckmu32/GoBankingCore/services"
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

	response := services.AddAccount(&account)

	s, _ := json.Marshal(response)

	/*
		s, _ := json.Marshal(struct {
			Code        int         `json:"code"`
			Description interface{} `json:"cuentas"`
			Errors      []string
		}{
			Code:        response.Code,
			Description: response.Description,
			Errors:      response.Errors,
		},
		)
	*/

	w.Header().Add("Content-Type", "application/json") //Must be before the http code assignation.
	w.WriteHeader(response.Code)
	w.Write(s)

}
