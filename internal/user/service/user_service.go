package service

import (
	"context"
	"dplatform/internal/user/domain"
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

func (s *userService) CreateUser(c context.Context, input *domain.User) (newUser *domain.User, err error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user, err := s.userRepository.CreateUser(ctx, input)
	if err != nil {
		return nil, err
	}

	return user, nil
}
