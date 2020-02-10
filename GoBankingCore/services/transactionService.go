package services

import "net/http"

// AddTransaction Let's us add a new transaction for an account or card.
func addTransaction(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Transaction added!"))
}
