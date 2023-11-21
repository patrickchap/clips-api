package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5/pgtype"
	mockdb "github.com/patrickchap/clipsapi/db/mock"
	db "github.com/patrickchap/clipsapi/db/sqlc"
	"github.com/patrickchap/clipsapi/util"
	"github.com/stretchr/testify/require"
)
type VideoWithLike struct {
	ID           int64            `json:"id"`
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	FileUrl      string           `json:"file_url"`
	ThumbnailUrl string           `json:"thumbnail_url"`
	UserID       string           `json:"user_id"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
}

func TestGetVideoAPI(t *testing.T){
	video := createRandomVideo()

	testCases := []struct{
		name string
		videoID int64
		buildStubs func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			videoID: video.ID,
			buildStubs: func(store *mockdb.MockStore){
				store.EXPECT().GetVideoWithLikes(gomock.Any(), gomock.Eq(video.ID)).Times(1).Return(video, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder){

				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchVideo(t, recorder.Body, video)
			},
		},
		//TODO: add more test cases
	}

	for i := range testCases{
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T){
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := NewServer(store) 
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/api/v1/video/%v", tc.videoID)
			request, err := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})

	}
}

func createRandomUser() db.User {
	user := db.User{
		ID: util.RandomInt(1, 100),
		Username: util.RandomUserName(),
		Email: util.RandomEmail(),
		Auth0UserID: util.RandomUserName() + "_Auth",
	}

	return user 
}

func createRandomVideo() db.GetVideoWithLikesRow {
	user := createRandomUser()

	video := db.GetVideoWithLikesRow{
		ID: util.RandomInt(1, 100),
		Title : util.RandomString(6),
		Description: util.RandomString(25),
		FileUrl: util.RandomString(10),
		UserID: user.Auth0UserID,
		ThumbnailUrl: "thumbnailurl",
		LikeCount: 0,
	}

	return video
}

func requireBodyMatchVideo(t *testing.T, body *bytes.Buffer, video db.GetVideoWithLikesRow){
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotVideo db.GetVideoWithLikesRow
	err = json.Unmarshal(data, &gotVideo)
	require.NoError(t, err)
	require.Equal(t, video, gotVideo)
}


