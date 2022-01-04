package service

import (
	"context"
	"dplatform/internal/room/domain"
	"time"
)

type roomService struct {
	roomRepository domain.RoomRepository
	timeout        time.Duration
}

func NewRoomService(roomRepository domain.RoomRepository) domain.RoomService {
	return &roomService{
		roomRepository,
		time.Duration(2) * time.Second,
	}
}

func (s *roomService) GetRoom(c context.Context, id int) (*domain.Room, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	result, err := s.roomRepository.GetRoom(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *roomService) CreateRoom(c context.Context, arg *domain.CreateRoomParams) (*domain.Room, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	room, err := s.roomRepository.CreateRoom(ctx, arg)
	if err != nil {
		return nil, err
	}

	return room, nil
}
