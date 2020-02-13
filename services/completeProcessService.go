package services

import (
	"net/http"
)

// CompleteProcess Serves the process of adding a new cardholder and creating it's account with the card. Responds to POST.
func CompleteProcess(w http.ResponseWriter, r *http.Request) {
	//Check that the card holder does not exits.
	//Check that the account does not exists.
	//Check that the card does not exists.
	//Create security service.
	w.Write([]byte("hello!"))
}
