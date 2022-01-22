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
	rows, err := r.DB.QueryxContext(ctx, ListSpacesQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []domain.Space{}
	for rows.Next() {
		var i domain.Space
		if err := rows.Scan(&i.ID, &i.Name, &i.Description, &i.Location, &i.Link); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *spaceDBRepository) CreateSpace(ctx context.Context, arg *domain.CreateSpaceParams) (*domain.Space, error) {
	lastInsertId := 0
	err := r.DB.QueryRowxContext(ctx, CreateSpaceQuery, arg.Name, arg.Description, arg.Location, util.RandomLink("davidhwang_ij"), arg.HostID, pq.Array(arg.Images), pq.Array(arg.Dates)).Scan(&lastInsertId)
	if err != nil {
		return nil, err
	}
	return r.GetSpace(ctx, lastInsertId)
}
