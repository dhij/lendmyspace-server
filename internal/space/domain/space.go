package domain

import (
	"context"
	"time"
)

type Space struct {
	ID          int64      `json:"id"`
	Name        *string    `json:"name" db:"name"`
	Description *string    `json:"description" db:"description"`
	Image       *string    `json:"image" db:"image"`
	HostID      int64      `json:"host_id" db:"host_id"`
	DateID      *int64     `json:"date_id" db:"date_id"`
	CreatedAt   time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt   *time.Time `json:"updatedAt" db:"updated_at"`
}

type CreateSpaceParams struct {
	ID          int64   `json:"id"`
	Name        *string `json:"name" db:"name"`
	Description *string `json:"description" db:"description"`
	Image       *string `json:"image" db:"image"`
	HostID      int64   `json:"host_id" db:"host_id"`
	DateID      *int64  `json:"date_id" db:"date_id"`
}

type SpaceService interface {
	GetSpace(ctx context.Context, id int) (*Space, error)
	CreateSpace(ctx context.Context, arg *CreateSpaceParams) (*Space, error)
}

type SpaceRepository interface {
	GetSpace(ctx context.Context, id int) (*Space, error)
	CreateSpace(ctx context.Context, arg *CreateSpaceParams) (*Space, error)
}
