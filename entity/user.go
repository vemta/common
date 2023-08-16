package entity

import "time"

type User struct {
	Username  string    `json:"username"`
	FullName  string    `json:"full_name"`
	Birthdate time.Time `json:"birthdate"`
}
