package entity

import "database/sql"

type Account struct {
	Id       int             `json:"id,omitempty"`
	Email    *sql.NullString `json:"email,omitempty"`
	Password *sql.NullString `json:"password,omitempty"`
	UserId   *sql.NullInt64  `json:"user_id,omitempty"`
}
