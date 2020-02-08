package models

type cardSecurity struct { //Encrypted.
	id         float64
	cardNumber string //PAN.
	expiryDate string
	startDate  string
	pin        int
}
