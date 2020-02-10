package models

// Transaction Represents the transaction that the cardholder made in a commerce.
type Transaction struct {
	ID              float64
	Description     string
	Date            string
	Time            string
	TransactionType string //plus or minus.
	CardNumber      string
	AccountNumber   string
	PossibleFraud   bool
	Recurring       bool
}
