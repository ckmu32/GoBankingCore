package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ckmu32/GoBankingCore/models"
	"github.com/ckmu32/GoBankingCore/services"
)

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
	//20 74
	services.AddAccount(&account)
}
