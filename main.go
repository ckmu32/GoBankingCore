package main

import (
	"net/http"

	accountServiceX "github.com/ckmu32/GoBankingCore/service"
)

func main() {
	http.HandleFunc("/", controller) //This URL responds with this method.
	http.HandleFunc("/account", accountServiceX.AddAccount)
	http.ListenAndServe(":8080", nil)
}

func controller(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		println("GET METHOD")
	} else if r.Method == http.MethodPost {
		println("POST METHOD")
	}
	w.Write([]byte("hello!"))
}
