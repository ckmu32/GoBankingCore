package services

import "net/http"

// AddAccount Let's us add a new account.
func AddAccount(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Account added!"))
}
