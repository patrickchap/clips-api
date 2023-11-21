package db

import (
	"context"
	"testing"

	"github.com/patrickchap/clipsapi/util"
	"github.com/stretchr/testify/require"
)


func createRandomVideo(t *testing.T) Video {
	user := createRandomUser(t)

	arg := CreateVideoParams{
		Title : util.RandomString(6),
		Description: util.RandomString(25),
		FileUrl: util.RandomString(10),
		UserID: user.Auth0UserID,
		ThumbnailUrl: "thumbnailurl",
	}

	video, err := testQueries.CreateVideo(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, video.ID)
	require.Equal(t, video.Title, arg.Title)
	require.Equal(t, video.Description, arg.Description)
	require.Equal(t, video.FileUrl, arg.FileUrl)
	require.Equal(t, video.UserID, arg.UserID)
	require.Equal(t, video.ThumbnailUrl, arg.ThumbnailUrl)
	return video
}

func TestCreateVideo(t *testing.T){
	createRandomVideo(t)
}



