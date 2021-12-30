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

func (r *userDBRepository) GetUser(ctx context.Context, id int) (*domain.User, error) {
	s := domain.User{}
	err := r.DB.GetContext(ctx, &s, GetUserQuery, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *userDBRepository) ListUsers(ctx context.Context, arg domain.ListUsersParams) ([]domain.User, error) {
	rows, err := r.DB.QueryxContext(ctx, ListUsersQuery, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domain.User
	for rows.Next() {
		var i domain.User
		if err := rows.Scan(&i.UserName, &i.FirstName, &i.LastName, &i.Email); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *userDBRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	lastInsertId := 0
	err := r.DB.QueryRowxContext(ctx, CreateUserQuery, user.UserName, user.FirstName, user.LastName, user.Email, user.Password).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}

	return r.GetUser(ctx, lastInsertId)
}
