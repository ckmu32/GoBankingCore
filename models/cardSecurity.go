package models

// CardSecurity Represents the sensible data of card that should be encrypted.
type CardSecurity struct { //Encrypted.
	ID               float64 `validate:"required"`
	CardNumber       string  `validate:"required,numeric,len=16"` //PAN.
	ExpiryDate       string  `validate:"required"`
	StartDate        string  //Set when the cardholder activate it,
	Pin              int     `validate:"len=4"`
	CreationDate     string  //Set automatically.
	CreationHour     string  //Set automatically.
	CreationUser     string  //Set automatically.
	ModificationDate string  //Set automatically.
	ModificationHour string  //Set automatically.
	ModificationUser string  //Set automatically.
}
