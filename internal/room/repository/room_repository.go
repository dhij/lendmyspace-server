package repository

import (
	"context"
	"database/sql"
	"dplatform/internal/room/domain"
	"errors"

	"github.com/jmoiron/sqlx"
)

type roomRepository struct {
	DB *sqlx.DB
}

func NewRoomRepository(db *sqlx.DB) domain.RoomRepository {
	return &roomRepository{
		DB: db,
	}
}

func (r *roomRepository) GetRoom(ctx context.Context, id int) (*domain.Room, error) {
	s := domain.Room{}
	err := r.DB.GetContext(ctx, &s, GetRoomQuery, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *roomRepository) CreateRoom(ctx context.Context, arg *domain.CreateRoomParams) (*domain.Room, error) {
	lastInsertId := 0
	err := r.DB.QueryRowxContext(ctx, CreateRoomQuery, arg.Name, arg.Description, arg.HostID, arg.ParticipantID, arg.MessageID).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}
	return r.GetRoom(ctx, lastInsertId)
}
