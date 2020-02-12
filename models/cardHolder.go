package models

// CardHolder Represents the person that owns an account and cards.
type CardHolder struct {
	ID               float64 `json:"id" validate:"required"`
	Name             string  `json:"name" validate:"required,max=50"`
	FatherName       string  `json:"father_name" validate:"required,max=20"`
	MotherName       string  `json:"mother_name" validate:"required,max=20"`
	DateOfBirth      string  `json:"date_of_birth" validate:"required"`
	Age              int     `json:"age" validate:"required"`
	Ssn              string  `json:"ssn" validate:"required,max=32"`
	Rfc              string  `json:"rfc" validate:"required,rfcVal"`
	Gender           string  `json:"gender" validate:"required,genVal"`
	Country          string  `json:"country" validate:"required,countryVal"`
	State            string  `json:"state" validate:"required,stateVal"`
	City             string  `json:"city" validate:"required,cityVal"`
	Address          string  `json:"address" validate:"required"`
	StreetNumber     int     `json:"street_number" validate:"required"`
	EducationLevel   string  `json:"education_level" validate:"required,educationVal"`
	Occupation       string  `json:"occupation" validate:"required"`
	CivilStatus      string  `json:"civil_status" validate:"required,civilStatusVal"`
	Email            string  `json:"email" validate:"required,email"`
	Cellphone        string  `json:"cellphone" validate:"required,max=12"`
	CreationDate     string  `json:"creation_date"`     //Set automatically.
	CreationHour     string  `json:"creation_hour"`     //Set automatically.
	CreationUser     string  `json:"creation_user"`     //Set automatically.
	ModificationDate string  `json:"mofication_date"`   //Set automatically.
	ModificationHour string  `json:"mofication_hour"`   //Set automatically.
	ModificationUser string  `json:"modification_user"` //Set automatically.
}
