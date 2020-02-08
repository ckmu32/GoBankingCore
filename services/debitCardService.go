package services

import "net/http"

// AddDebitCard Let's us add a new debit card.
func addDebitCard(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Account added!"))
}
