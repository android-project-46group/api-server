package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/android-project-46group/api-server/db"
	mockdb "github.com/android-project-46group/api-server/db/mock"
	models "github.com/android-project-46group/api-server/db/my_models"
	"github.com/android-project-46group/api-server/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAllFormationsAPI(t *testing.T) {

	groupName := "nogizaka"
	key := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"

	testCases := []struct {
		name          string
		url           string
		buildStubs    func(store *mockdb.MockQuerier)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			url:  fmt.Sprintf("/formations?gn=%s&key=%s", groupName, key),
			buildStubs: func(store *mockdb.MockQuerier) {
				store.EXPECT().
					GetAllFormations(groupName).
					Times(1).
					Return([]db.PositionSongsBind{}, nil)
				store.EXPECT().
					FindApiKeyByName(key).
					Times(1).
					Return(nil, nil)
				store.EXPECT().
					FindGroupByName(groupName).
					Times(1).
					Return(&models.Group{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "InvalidAPIKey",
			url:  fmt.Sprintf("/formations?gn=%s&key=%s", groupName, "invalid_key"),
			buildStubs: func(store *mockdb.MockQuerier) {
				store.EXPECT().
					GetAllFormations(gomock.Any()).
					Times(0)
				store.EXPECT().
					FindApiKeyByName("invalid_key").
					Times(1).
					Return(nil, sql.ErrNoRows)
				store.EXPECT().
					FindGroupByName(gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "NotExistingGroup",
			url:  fmt.Sprintf("/formations?gn=%s&key=%s", "not_existing_group", key),
			buildStubs: func(store *mockdb.MockQuerier) {
				store.EXPECT().
					GetAllFormations(gomock.Any()).
					Times(0)
				store.EXPECT().
					FindApiKeyByName(key).
					Times(1).
					Return(nil, nil)
				store.EXPECT().
					FindGroupByName("not_existing_group").
					Times(1).
					Return(&models.Group{}, fmt.Errorf("Failed to FindGroupByName: %w", sql.ErrNoRows))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InternalDBError",
			url:  fmt.Sprintf("/formations?gn=%s&key=%s", groupName, key),
			buildStubs: func(store *mockdb.MockQuerier) {
				store.EXPECT().
					GetAllFormations(groupName).
					Times(1).
					Return(nil, errors.New("internal server error"))
				store.EXPECT().
					FindApiKeyByName(key).
					Times(1).
					Return(nil, nil)
				store.EXPECT().
					FindGroupByName(groupName).
					Times(1).
					Return(&models.Group{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			config := util.Config{
				IsCGI:         false,
				ServerAddress: "0.0.0.0:8080",
			}
			querier := mockdb.NewMockQuerier(ctrl)
			tc.buildStubs(querier)
			mathcer := util.NewMatcher()
			server, err := NewServer(config, querier, mathcer)
			require.NoError(t, err)
			recorder := httptest.NewRecorder()

			request, err := http.NewRequest(http.MethodGet, tc.url, nil)
			require.NoError(t, err)

			require.NotNil(t, server.router)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}
