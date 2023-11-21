package db

import (
	"context"
	"testing"
	"time"

	"github.com/patrickchap/clipsapi/util"
	"github.com/stretchr/testify/require"
)
func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username: util.RandomUserName(),
		Email: util.RandomEmail(),
		Auth0UserID: util.RandomUserName() + "_Auth",
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Auth0UserID, user.Auth0UserID)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt) 

	return user 
}
func TestCreateUser(t *testing.T){
	createRandomUser(t)
}


func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Auth0UserID, user2.Auth0UserID)
	require.WithinDuration(t, user1.CreatedAt.Time, user2.CreatedAt.Time, time.Second)
}

