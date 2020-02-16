package models

// Account The principal data that an account has. The account the represents the infotmation that stores the balance of the card holder.
type Account struct {
	Number           string  `json:"number" validate:"numeric,required,len=15"`
	CardHolderID     float64 `json:"card_holder_id" validate:"required"`                  //10
	Type             string  `json:"type" validate:"required,accountTypeVal"`             //10
	ExternalAccount  string  `json:"external_account" validate:"numeric,required,len=25"` //The account visible to external people for deposits or other things.
	Level            string  `json:"level" validate:"required,accountLevelVal"`           //10
	AvailableBalance float64 `json:"available_balance"`                                   //The balance before taxes, approved payments, etc. Is confirmed. Can be zero.
	RealBalance      float64 `json:"real_balance"`                                        //The balance that the cardHolder can use to pay. Can be zero.
	Status           string  `json:"status" validate:"required,accountStatusVal"`         //10
	CreationDate     string  `json:"creation_date"`                                       //Set automatically.
	CreationHour     string  `json:"creation_hour"`                                       //Set automatically.
	CreationUser     string  `json:"creation_user"`                                       //Set automatically.
	ModificationDate string  `json:"mofication_date"`                                     //Set automatically.
	ModificationHour string  `json:"mofication_hour"`                                     //Set automatically.
	ModificationUser string  `json:"modification_user"`                                   //Set automatically.
}
