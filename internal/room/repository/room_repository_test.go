package repository

import (
	"context"
	"dplatform/internal/room/domain"
	"dplatform/util"
	"testing"

	"github.com/stretchr/testify/require"
)

var roomRepository domain.RoomRepository

func createRandomRoom(t *testing.T) *domain.Room {
	rs := util.RandomString(6)
	description := util.RandomString(20)
	var pid int64 = 3

	arg := domain.CreateRoomParams{
		Name:          &rs,
		Description:   &description,
		HostID:        1,
		ParticipantID: &pid,
	}

	roomRepository = NewRoomRepository(dbSQLX.GetDB())
	room, err := roomRepository.CreateRoom(context.Background(), &arg)
	require.NoError(t, err)
	require.NotEmpty(t, room)

	require.Equal(t, arg.Name, room.Name)
	require.Equal(t, arg.Description, room.Description)
	require.Equal(t, arg.HostID, room.HostID)
	require.Equal(t, arg.ParticipantID, room.ParticipantID)

	require.NotZero(t, room.ID)
	return room
}

func TestCreateRoom(t *testing.T) {
	createRandomRoom(t)
}

func TestGetRoom(t *testing.T) {
	room1 := createRandomRoom(t)

	room2, err := roomRepository.GetRoom(context.Background(), int(room1.ID))
	require.NoError(t, err)
	require.NotEmpty(t, room2)

	require.Equal(t, room1.ID, room2.ID)
	require.Equal(t, room1.Name, room2.Name)
	require.Equal(t, room1.Description, room2.Description)
	require.Equal(t, room1.HostID, room2.HostID)
	require.Equal(t, room1.ParticipantID, room2.ParticipantID)
}
