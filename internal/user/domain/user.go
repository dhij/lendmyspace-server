package domain

import (
	"context"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	UserName  string    `json:"user_name" db:"user_name"`
	FirstName string    `json:"first_name" db:"first_name"`
	LastName  string    `json:"last_name" db:"last_name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	LastLogin time.Time `json:"last_login" db:"last_login"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (newUser *User, err error)
	FindByID(ctx context.Context, id int) (user *User, err error)
}

type UserService interface {
	CreateUser(ctx context.Context, user *User) (newUser *User, err error)
}
