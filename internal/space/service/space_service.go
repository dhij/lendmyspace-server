package service

import (
	"context"
	"lendmyspace-server/internal/space/domain"

	"time"
)

type spaceService struct {
	spaceRepository domain.SpaceRepository
	timeout         time.Duration
}

func NewSpaceService(spaceRepository domain.SpaceRepository) domain.SpaceService {
	return &spaceService{
		spaceRepository,
		time.Duration(2) * time.Second,
	}
}

func (s *spaceService) GetSpace(c context.Context, id int) (*domain.Space, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	result, err := s.spaceRepository.GetSpace(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *spaceService) ListSpaces(c context.Context) ([]domain.Space, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	result, err := s.spaceRepository.ListSpaces(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *spaceService) CreateSpace(c context.Context, arg *domain.CreateSpaceParams) (*domain.Space, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	space, err := s.spaceRepository.CreateSpace(ctx, arg)
	if err != nil {
		return nil, err
	}

	return space, nil
}
