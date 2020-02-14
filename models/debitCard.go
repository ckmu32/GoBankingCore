package models

// DebitCard Represents the basic product that a card holder can obtain.
type DebitCard struct {
	ID                   float64 `json:"id" validate:"required"`
	AccountNumber        string  `json:"account_number" validate:"numeric,required,len=15"` //The account related to this card.
	CardType             string  `json:"card_type" validate:"required,debitCardTypeVal"`
	Status               string  `json:"status" validate:"required,debitCardStatusVal"`
	Cancelled            bool    `json:"cancelled" validate:"required"`
	Locked               bool    `json:"locked" validate:"required"`
	EmergencyReplacement bool    `json:"emergency_replacement" validate:"required"`
	SecurityID           string  `json:"security_id"`       //Set automatically. Sensible data of the card separated for PCI (Start of it).
	CreationDate         string  `json:"creation_date"`     //Set automatically.
	CreationHour         string  `json:"creation_hour"`     //Set automatically.
	CreationUser         string  `json:"creation_user"`     //Set automatically.
	ModificationDate     string  `json:"mofication_date"`   //Set automatically.
	ModificationHour     string  `json:"mofication_hour"`   //Set automatically.
	ModificationUser     string  `json:"modification_user"` //Set automatically.
}
