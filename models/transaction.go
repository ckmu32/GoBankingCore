package models

type transaction struct {
	id              float64
	description     string
	date            string
	time            string
	transactionType string //plus or minus.
	cardNumber      string
	accountNumber   string
	possibleFraud   bool
	recurring       bool
}
