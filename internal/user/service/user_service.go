package service

import (
	"context"
	"lendmyspace-server/internal/user/domain"
	"time"
)

type userService struct {
	userRepository domain.UserRepository
	timeout        time.Duration
}

func NewUserSerivce(userRepository domain.UserRepository) domain.UserService {
	return &userService{
		userRepository,
		time.Duration(2) * time.Second,
	}
}

func (s *userService) GetUser(c context.Context, id int) (user *domain.UserInfo, err error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	result, err := s.userRepository.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *userService) ListUsers(c context.Context, arg domain.ListUsersParams) ([]domain.UserInfo, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	result, err := s.userRepository.ListUsers(ctx, arg)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *userService) CreateUser(c context.Context, arg *domain.User) (newUser *domain.UserInfo, err error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user, err := s.userRepository.CreateUser(ctx, arg)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) UpdateUser(c context.Context, arg domain.UpdateUserParams) (*domain.UserInfo, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()
	user, err := s.userRepository.UpdateUser(ctx, arg)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) DeleteUser(c context.Context, id int) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()
	err := s.userRepository.DeleteUser(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
