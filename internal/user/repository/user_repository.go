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

func (r *userDBRepository) GetUser(ctx context.Context, id int) (*domain.UserInfo, error) {
	s := domain.UserInfo{}
	err := r.DB.GetContext(ctx, &s, GetUserQuery, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *userDBRepository) ListUsers(ctx context.Context, arg domain.ListUsersParams) ([]domain.UserInfo, error) {
	rows, err := r.DB.QueryxContext(ctx, ListUsersQuery, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []domain.UserInfo
	for rows.Next() {
		var i domain.UserInfo
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

func (r *userDBRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.UserInfo, error) {
	lastInsertId := 0
	err := r.DB.QueryRowxContext(ctx, CreateUserQuery, user.UserName, user.FirstName, user.LastName, user.Email, user.Password).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}

	return r.GetUser(ctx, lastInsertId)
}

func (r *userDBRepository) UpdateUser(ctx context.Context, arg domain.UpdateUserParams) (*domain.UserInfo, error) {
	lastInsertId := 0
	err := r.DB.QueryRowxContext(ctx, UpdateUserQuery, arg.ID, arg.UserName).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}

	return r.GetUser(ctx, lastInsertId)
}

func (r *userDBRepository) DeleteUser(ctx context.Context, id int) error {
	_, err := r.DB.ExecContext(ctx, DeleteUserQuery, id)
	return err
}
