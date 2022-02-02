package repository

import (
	"context"
	"lendmyspace-server/internal/space/domain"
	"lendmyspace-server/util"
	"testing"

	"github.com/stretchr/testify/require"
)

var spaceRepository domain.SpaceRepository

func createRandomSpace(t *testing.T) *domain.Space {
	rs := util.RandomString(6)
	description := util.RandomString(20)

	arg := domain.CreateSpaceParams{
		Name:        &rs,
		Description: &description,
		HostID:      1,
		Images:      util.RandomSlice(3),
		Dates:       util.RandomSlice(3),
	}

	spaceRepository = NewSpaceRepository(dbSQLX.GetDB())
	space, err := spaceRepository.CreateSpace(context.Background(), &arg)
	require.NoError(t, err)
	require.NotEmpty(t, space)

	require.Equal(t, arg.Name, space.Name)
	require.Equal(t, arg.Description, space.Description)
	require.Equal(t, arg.HostID, space.HostID)
	require.Equal(t, arg.Images, []string(space.Images))
	require.Equal(t, arg.Dates, []string(space.Dates))

	require.NotZero(t, space.ID)
	return space
}

func TestCreateSpace(t *testing.T) {
	createRandomSpace(t)
}

func TestGetSpace(t *testing.T) {
	space1 := createRandomSpace(t)

	space2, err := spaceRepository.GetSpace(context.Background(), int(space1.ID))
	require.NoError(t, err)
	require.NotEmpty(t, space2)

	require.Equal(t, space1.ID, space2.ID)
	require.Equal(t, space1.Name, space2.Name)
	require.Equal(t, space1.Description, space2.Description)
	require.Equal(t, space1.HostID, space2.HostID)
	require.Equal(t, space1.Images, space2.Images)
	require.Equal(t, space1.Dates, space2.Dates)
}
