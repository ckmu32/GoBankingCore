package services

import "net/http"

// AddCardHolder Let's us add a new card holder.
func addCardHolder(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Card Holder added!"))
}
