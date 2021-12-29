package repository

import (
	"context"
	"database/sql"
	"dplatform/internal/user/domain"
	"errors"

	"github.com/jmoiron/sqlx"
)

type userDBRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) domain.UserRepository {
	return &userDBRepository{
		db,
	}
}

func (r *userDBRepository) FindByID(ctx context.Context, id int) (*domain.User, error) {
	s := domain.User{}
	err := r.DB.GetContext(ctx, &s, FindByIDQuery, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *userDBRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	lastInsertId := 0
	err := r.DB.QueryRowxContext(ctx, CreateUserQuery, user.UserName, user.FirstName, user.LastName, user.Email, user.Password).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}

	return r.FindByID(ctx, lastInsertId)
}
