package repository

import (
	"context"
	"database/sql"
	"errors"
	"lendmyspace-server/internal/user/domain"

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

func (r *userDBRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	s := domain.User{}
	err := r.DB.GetContext(ctx, &s, GetUserByEmailQuery, email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *userDBRepository) ListUsers(ctx context.Context) ([]domain.UserInfo, error) {
	users := []domain.UserInfo{}
	err := r.DB.SelectContext(ctx, &users, ListUsersQuery)
	if err != nil {
		return nil, err
	}

	return users, nil
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
