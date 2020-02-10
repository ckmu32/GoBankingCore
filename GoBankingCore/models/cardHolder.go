package models

// CardHolder Represents the person that owns an account and cards.
type CardHolder struct {
	ID             float64
	Name           string
	FatherName     string
	MotherName     string
	DateOfBirth    string
	Age            int
	Ssn            string
	Rfc            string
	Gender         string
	Country        string
	State          string
	City           string
	Address        string
	StreetNumber   int
	EducationLevel string
	Occupation     string
	CivilStatus    string
	Email          string
	Cellphone      string
}
