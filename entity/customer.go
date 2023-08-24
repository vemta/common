package entity

import "time"

type User struct {
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	Birthdate time.Time `json:"birthdate"`
}
