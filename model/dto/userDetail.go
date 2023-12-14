package dto

type UserDetail struct {
	Id          int    `json:"id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Gender      string `json:"gender,omitempty"`
	BirthDate   string `json:"birth_date,omitempty"`
	AddressId   int    `json:"address_id,omitempty"`
	EducationId int    `json:"education_id,omitempty"`
}
