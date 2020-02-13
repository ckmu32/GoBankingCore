package models

// Transaction Represents the transaction that the cardholder made in a commerce.
type Transaction struct {
	ID            float64 `json:"id"`
	Description   string  `json:"description" validate:"required"`
	Date          string  `json:"date" validate:"required"`
	Time          string  `json:"time" validate:"required"`
	Type          string  `json:"type" validate:"transactionTypeVal"` //plus or minus.
	CardNumber    string  `json:"card_number" validate:"required,numeric,len=16"`
	Fees          float64 `json:"fees" validate:"numeric"`
	AccountNumber string  `json:"accountNumber" validate:"numeric,required,len=15"`
	PossibleFraud bool    `json:"posible_fraud" validate:"required"`
	Recurring     bool    `json:"recurring" validate:"required"`
	CreationDate  string  `json:"creation_date"` //Set automatically.
	CreationHour  string  `json:"creation_hour"` //Set automatically.
}
