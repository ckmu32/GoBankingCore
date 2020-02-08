package models

// Account The principal data that an account has. The account the represents the infotmation that stores the balance of the card holder.
type Account struct {
	AccountNumber    string
	CardHolderID     float64
	AccountType      string
	ExternalAccount  string //The account visible to external people for deposits or other things.
	AccountLevel     string
	AvailableBalance float64 //The balance before taxes, approved payments, etc. Is confirmed.
	RealBalance      float64 //The balance that the cardHolder can use to pay.
	Status           string
}
