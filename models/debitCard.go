package models

type debitCard struct {
	id            float64
	accountNumber string //The account related to this card.
	cardType      string
	status        string
	cancelled     bool
	locked        bool
	securityID    string //Sensible data of the card separated for PCI (Start of it).
}
