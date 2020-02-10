package models

// CardSecurity Represents the sensible data of card that should be encrypted.
type CardSecurity struct { //Encrypted.
	ID         float64
	CardNumber string //PAN.
	ExpiryDate string
	StartDate  string
	Pin        int
}
