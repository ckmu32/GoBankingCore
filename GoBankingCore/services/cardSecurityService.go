package services

import "net/http"

// AddCardSecurity Let's us add a new security for a card (Sensible data).
func addCardSecurity(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Account added!"))
}
