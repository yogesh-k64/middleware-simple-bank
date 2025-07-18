package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yogesh-k64/middleware-simple-bank/utils"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:       utils.RandomOwner(),
		HashedPassword: "1234",
		FullName:       utils.RandomOwner(),
		Email:          utils.RandomEmail(),
	}
	// Create a new user
	user, err := testQueries.CreateUser(context.Background(), arg)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.NotZero(t, user.CreatedAt)
	require.True(t, user.PasswordChangedAt.IsZero())

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)
	testUser, err := testQueries.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, testUser)

	require.Equal(t, user.Username, testUser.Username)
	require.Equal(t, user.HashedPassword, testUser.HashedPassword)
	require.Equal(t, user.FullName, testUser.FullName)
	require.Equal(t, user.Email, testUser.Email)
	require.WithinDuration(t, user.CreatedAt, testUser.CreatedAt, time.Second)
}
