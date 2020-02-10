package models

// Account The principal data that an account has. The account the represents the infotmation that stores the balance of the card holder.
type Account struct {
	Number           string  `json:"number"`
	CardHolderID     float64 `json:"card_holder_id"`
	Type             string  `json:"type"`
	ExternalAccount  string  `json:"external_account"` //The account visible to external people for deposits or other things.
	Level            string  `json:"level"`
	AvailableBalance float64 `json:"available_balance"` //The balance before taxes, approved payments, etc. Is confirmed.
	RealBalance      float64 `json:"real_balance"`      //The balance that the cardHolder can use to pay.
	Status           string  `json:"status"`
}
