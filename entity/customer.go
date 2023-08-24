package entity

import "time"

type Customer struct {
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	Birthdate time.Time `json:"birthdate"`
}
