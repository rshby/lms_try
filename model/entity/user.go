package entity

import "database/sql"

type User struct {
	Id          int             `json:"id,omitempty"`
	FirstName   *sql.NullString `json:"first_name,omitempty"`
	LastName    *sql.NullString `json:"last_name,omitempty"`
	Gender      *sql.NullString `json:"gender,omitempty"`
	BirthDate   *sql.NullTime   `json:"birthdate,omitempty"`
	AddressId   *sql.NullInt64  `json:"address_id,omitempty"`
	EducationId *sql.NullInt64  `json:"education_id,omitempty"`
}
