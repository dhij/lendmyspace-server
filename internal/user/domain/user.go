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

type UserInfo struct {
	ID        int64  `json:"id"`
	UserName  string `json:"user_name" db:"user_name"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
	Email     string `json:"email" db:"email"`
}

type UpdateUserParams struct {
	ID       int64  `json:"id"`
	UserName string `json:"user_name"`
}

type LoginUserRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginUserResponse struct {
	AccessToken string    `json:"access_token"`
	User        *UserInfo `json:"user"`
}

type UserRepository interface {
	GetUser(ctx context.Context, id int) (*UserInfo, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	ListUsers(ctx context.Context) ([]UserInfo, error)
	CreateUser(ctx context.Context, arg *User) (*UserInfo, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (*UserInfo, error)
	DeleteUser(ctx context.Context, id int) error
}

type UserService interface {
	GetUser(ctx context.Context, id int) (*UserInfo, error)
	GetUserByEmail(c context.Context, email string) (user *User, err error)
	ListUsers(ctx context.Context) ([]UserInfo, error)
	CreateUser(ctx context.Context, arg *User) (*UserInfo, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (*UserInfo, error)
	DeleteUser(ctx context.Context, id int) error
}
