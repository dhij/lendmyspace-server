package domain

import (
	"context"
	"time"

	"github.com/lib/pq"
)

type Space struct {
	ID          int64          `json:"id"`
	Name        *string        `json:"name" db:"name"`
	Description *string        `json:"description" db:"description"`
	Location    *string        `json:"location" db:"location"`
	Link        *string        `json:"link" db:"link"`
	HostID      int64          `json:"host_id" db:"host_id"`
	Images      pq.StringArray `json:"images" db:"images"`
	Dates       pq.StringArray `json:"dates" db:"dates"`
	CreatedAt   time.Time      `json:"createdAt" db:"created_at"`
	UpdatedAt   *time.Time     `json:"updatedAt" db:"updated_at"`
}

type CreateSpaceParams struct {
	ID          int64    `json:"id"`
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Location    *string  `json:"location"`
	HostID      int64    `json:"host_id"`
	Images      []string `json:"images"`
	Dates       []string `json:"dates"`
}

type SpaceService interface {
	GetSpace(ctx context.Context, id int) (*Space, error)
	ListSpaces(ctx context.Context) ([]Space, error)
	CreateSpace(ctx context.Context, arg *CreateSpaceParams) (*Space, error)
}

type SpaceRepository interface {
	GetSpace(ctx context.Context, id int) (*Space, error)
	ListSpaces(ctx context.Context) ([]Space, error)
	CreateSpace(ctx context.Context, arg *CreateSpaceParams) (*Space, error)
}
