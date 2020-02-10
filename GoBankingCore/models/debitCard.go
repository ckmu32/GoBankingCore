package models

// DebitCard Represents the basic product that a card holder can obtain.
type DebitCard struct {
	ID            float64
	AccountNumber string //The account related to this card.
	CardType      string
	Status        string
	Cancelled     bool
	Locked        bool
	SecurityID    string //Sensible data of the card separated for PCI (Start of it).
}
