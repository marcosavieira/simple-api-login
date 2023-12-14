package db

import (
	"context"
	"testing"

	"github.com/marcosavieira/simple-api-login/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: util.RandomString(6),
		Password: util.RandomString(8),
		Email:    util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.Email, user.Email)
	require.NotEmpty(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Email, user2.Email)
	require.NotEmpty(t, user2.CreatedAt)
}
func TestGetUserById(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserById(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Password, user2.Password)
	require.Equal(t, user1.Email, user2.Email)
	require.NotEmpty(t, user2.CreatedAt)
}

func TestGetUsersValidUsername(t *testing.T) {
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	user3 := createRandomUser(t)

	listUsers1, err1 := testQueries.GetUsersByUsername(context.Background(), user1.Username)
	listUsers2, err2 := testQueries.GetUsersByUsername(context.Background(), user2.Username)
	listUsers3, err3 := testQueries.GetUsersByUsername(context.Background(), user3.Username)
	listUsersRandom, err := testQueries.GetUsersByUsername(context.Background(), "a")

	require.NoError(t, err)
	require.NoError(t, err1)
	require.NoError(t, err2)
	require.NoError(t, err3)
	require.NotEmpty(t, user1)
	require.NotEmpty(t, user2)
	require.NotEmpty(t, user3)
	require.NotEmpty(t, listUsers1)
	require.NotEmpty(t, listUsers2)
	require.NotEmpty(t, listUsers3)
	require.NotEmpty(t, listUsersRandom)
	require.ElementsMatch(t, []string{user1.Username}, listUsers1)
	require.ElementsMatch(t, []string{user2.Username}, listUsers2)
	require.ElementsMatch(t, []string{user3.Username}, listUsers3)

}

func TestGetUsersInvalidUsername(t *testing.T) {
	listUsers, err := testQueries.GetUsersByUsername(context.Background(), "nonexistent")

	require.NoError(t, err)
	require.Empty(t, listUsers)
}

func TestGetUsersCorrectUser(t *testing.T) {
	user := createRandomUser(t)

	listUsers, err := testQueries.GetUsersByUsername(context.Background(), user.Username)

	require.NoError(t, err)
	require.NotEmpty(t, listUsers)
	require.ElementsMatch(t, []string{user.Username}, listUsers)
}

func TestGetUsersContextCancelled(t *testing.T) {
	user := createRandomUser(t)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	listUsers, err := testQueries.GetUsersByUsername(ctx, user.Username)

	require.Error(t, err)
	require.Empty(t, listUsers)
}

func TestGetUsersQueryFail(t *testing.T) {
	// Simulate a query failure by passing an empty username
	listUsers, err := testQueries.GetUsersByUsername(context.Background(), "1")

	require.NoError(t, err)
	require.Empty(t, listUsers)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)
	arg := UpdateUserParams{
		Username: util.RandomString(6),
		Password: util.RandomString(8),
		Email:    util.RandomEmail(),
		ID:       user1.ID,
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, arg.Username, user2.Username)
	require.Equal(t, arg.Password, user2.Password)
	require.Equal(t, arg.Email, user2.Email)
	require.NotEmpty(t, user2.CreatedAt)
}

func TestDeleteUserID(t *testing.T) {
	user1 := createRandomUser(t)
	err := testQueries.DeleteUserID(context.Background(), user1.ID)
	require.NoError(t, err)
}

func TestDeleteUserUsername(t *testing.T) {
	user1 := createRandomUser(t)
	err := testQueries.DeleteUserUsername(context.Background(), user1.Username)
	require.NoError(t, err)
}
