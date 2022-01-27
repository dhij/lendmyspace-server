package repository

import (
	"context"
	"database/sql"
	"errors"
	"lendmyspace-server/internal/space/domain"
	"lendmyspace-server/util"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type spaceDBRepository struct {
	DB *sqlx.DB
}

func NewSpaceRepository(db *sqlx.DB) domain.SpaceRepository {
	return &spaceDBRepository{
		db,
	}
}

func (r *spaceDBRepository) GetSpace(ctx context.Context, id int) (*domain.Space, error) {
	s := domain.Space{}
	err := r.DB.GetContext(ctx, &s, GetSpaceQuery, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *spaceDBRepository) ListSpaces(ctx context.Context) ([]domain.Space, error) {
	spaces := []domain.Space{}
	err := r.DB.SelectContext(ctx, &spaces, ListSpacesQuery)
	if err != nil {
		return nil, err
	}

	return spaces, nil
}

func (r *spaceDBRepository) CreateSpace(ctx context.Context, arg *domain.CreateSpaceParams) (*domain.Space, error) {
	lastInsertId := 0
	err := r.DB.QueryRowxContext(ctx, CreateSpaceQuery, arg.Name, arg.Description, arg.Location, util.RandomLink("davidhwang_ij"), arg.HostID, pq.Array(arg.Images), pq.Array(arg.Dates)).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}
	return r.GetSpace(ctx, lastInsertId)
}
