package model

import "time"

type User struct {
	ID        int64     `json:"id"`
	UserName  string    `json:"user_name" db:"user_name"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	LastLogin time.Time `json:"last_login" db:"last_login"`
}
