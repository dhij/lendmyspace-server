package domain

import (
	"context"
	"time"
)

type Room struct {
	ID            int64      `json:"id"`
	Name          *string    `json:"name" db:"name"`
	Description   *string    `json:"description" db:"description"`
	HostID        int64      `json:"host_id" db:"host_id"`
	ParticipantID *int64     `json:"participant_id" db:"participant_id"`
	MessageID     *int64     `json:"message_id" db:"message_id"`
	CreatedAt     time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt     *time.Time `json:"updatedAt" db:"updated_at"`
}

type CreateRoomParams struct {
	Name          *string `json:"name" db:"name"`
	Description   *string `json:"description" db:"description"`
	HostID        int64   `json:"host_id" db:"host_id"`
	ParticipantID *int64  `json:"participant_id" db:"participant_id"`
	MessageID     *int64  `json:"message_id" db:"message_id"`
}

type RoomService interface {
	GetRoom(ctx context.Context, id int) (*Room, error)
	CreateRoom(ctx context.Context, arg *CreateRoomParams) (*Room, error)
}

type RoomRepository interface {
	GetRoom(ctx context.Context, id int) (*Room, error)
	CreateRoom(ctx context.Context, arg *CreateRoomParams) (*Room, error)
}
