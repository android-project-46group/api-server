package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	mockdb "github.com/android-project-46group/api-server/db/mock"
	models "github.com/android-project-46group/api-server/db/my_models"
	"github.com/android-project-46group/api-server/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAllGroupsAPI(t *testing.T) {

	key := "valid_api_key"

	testCases := []struct {
		name          string
		url           string
		buildStubs    func(querier *mockdb.MockQuerier)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			url:  fmt.Sprintf("/groups?key=%s", key),
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetAllGroups(gomock.Any()).
					Times(1).
					Return([]*models.Group{}, nil)
				querier.EXPECT().
					FindApiKeyByName(gomock.Any(), key).
					Times(1).
					Return(nil, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				require.NotNil(t, recorder.Body)
			},
		},
		{
			name: "NoAPIKeyInQueryParameter",
			url:  fmt.Sprintf("/groups"),
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetAllGroups(gomock.Any()).
					Times(0)
				querier.EXPECT().
					FindApiKeyByName(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InvalidAPIKey",
			url:  fmt.Sprintf("/groups?key=%s", "invalid_api_key"),
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetAllGroups(gomock.Any()).
					Times(0).
					Return([]*models.Group{}, nil)
				querier.EXPECT().
					FindApiKeyByName(gomock.Any(), "invalid_api_key").
					Times(1).
					Return(nil, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InternalDBErrorWhenReadingAPIKey",
			url:  fmt.Sprintf("/groups?key=%s", key),
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetAllGroups(gomock.Any()).
					Times(0).
					Return([]*models.Group{}, nil)
				querier.EXPECT().
					FindApiKeyByName(gomock.Any(), key).
					Times(1).
					Return(nil, sql.ErrConnDone)

			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InternalDBError",
			url:  fmt.Sprintf("/groups?key=%s", key),
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetAllGroups(gomock.Any()).
					Times(1).
					Return(nil, errors.New("internal server error"))
				querier.EXPECT().
					FindApiKeyByName(gomock.Any(), key).
					Times(1).
					Return(nil, nil)
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
			matcher := util.NewMatcher()
			logger, _, _ := util.NewStandardLogger("go-test", "api-saka", os.Stdout)

			server, err := NewServer(config, querier, matcher, logger, &mockGrpcClient{})
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
