package repository

import (
	"context"
	"database/sql"
	"lendmyspace-server/internal/user/domain"
	"lendmyspace-server/util"
	"testing"

	"github.com/stretchr/testify/require"
)

var userRepository domain.UserRepository

func createRandomUser(t *testing.T) *domain.UserInfo {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	arg := domain.User{
		UserName:  util.RandomUser(),
		FirstName: util.RandomFirstName(),
		LastName:  util.RandomLastName(),
		Email:     util.RandomEmail(),
		Password:  hashedPassword,
	}

	userRepository = NewUserRepository(dbSQLX.GetDB())
	newUser, err := userRepository.CreateUser(context.Background(), &arg)
	require.NoError(t, err)
	require.NotEmpty(t, newUser)

	require.Equal(t, arg.UserName, newUser.UserName)
	require.Equal(t, arg.FirstName, newUser.FirstName)
	require.Equal(t, arg.LastName, newUser.LastName)
	require.Equal(t, arg.Email, newUser.Email)

	require.NotZero(t, newUser.ID)
	return newUser
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)

	user2, err := userRepository.GetUser(context.Background(), int(user1.ID))
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.UserName, user2.UserName)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.Email, user2.Email)
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	users, err := userRepository.ListUsers(context.Background())
	require.NoError(t, err)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	arg := domain.UpdateUserParams{
		ID:       user1.ID,
		UserName: util.RandomUser(),
	}

	user2, err := userRepository.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, arg.UserName, user2.UserName)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.Email, user2.Email)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)

	err := userRepository.DeleteUser(context.Background(), int(user1.ID))
	require.NoError(t, err)

	user2, err := userRepository.GetUser(context.Background(), int(user1.ID))
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}
