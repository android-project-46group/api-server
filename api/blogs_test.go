package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/android-project-46group/api-server/db"
	mockdb "github.com/android-project-46group/api-server/db/mock"
	models "github.com/android-project-46group/api-server/db/my_models"
	"github.com/android-project-46group/api-server/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAllBlogsAPI(t *testing.T) {

	groupName := "nogizaka"
	key := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"

	testCases := []struct {
		name          string
		url           string
		buildStubs    func(querier *mockdb.MockQuerier)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			url:  fmt.Sprintf("/blogs?gn=%s&key=%s", groupName, key),
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetAllBlogs(gomock.Any(), groupName).
					Times(1).
					Return([]db.MemberBlogBind{}, nil)
				querier.EXPECT().
					FindApiKeyByName(gomock.Any(), key).
					Times(1).
					Return(nil, nil)
				querier.EXPECT().
					FindGroupByName(gomock.Any(), groupName).
					Times(1).
					Return(&models.Group{}, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				require.NotNil(t, recorder.Body)
			},
		},
		{
			name: "NoAPIKeyInQueryParameter",
			url:  fmt.Sprintf("/blogs?gn=%s", groupName),
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetAllBlogs(gomock.Any(), gomock.Any()).
					Times(0)
				querier.EXPECT().
					FindApiKeyByName(gomock.Any(), gomock.Any()).
					Times(0)
				querier.EXPECT().
					FindGroupByName(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name: "InvalidAPIKey",
			url:  fmt.Sprintf("/blogs?gn=%s&key=%s", groupName, "invalid_api_key"),
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetAllBlogs(gomock.Any(), gomock.Any()).
					Times(0).
					Return([]db.MemberBlogBind{}, nil)
				querier.EXPECT().
					FindApiKeyByName(gomock.Any(), "invalid_api_key").
					Times(1).
					Return(nil, sql.ErrNoRows)
				querier.EXPECT().
					FindGroupByName(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name: "InternalDBErrorWhenReadingAPIKey",
			url:  fmt.Sprintf("/blogs?gn=%s&key=%s", groupName, key),
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetAllBlogs(gomock.Any(), gomock.Any()).
					Times(0).
					Return([]db.MemberBlogBind{}, nil)
				querier.EXPECT().
					FindApiKeyByName(gomock.Any(), key).
					Times(1).
					Return(nil, sql.ErrConnDone)
				querier.EXPECT().
					FindGroupByName(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "NoGroupNameInQueryParameter",
			url:  fmt.Sprintf("/blogs?key=%s", key),
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetAllBlogs(gomock.Any(), "").
					Times(1).
					Return([]db.MemberBlogBind{}, nil)
				querier.EXPECT().
					FindApiKeyByName(gomock.Any(), key).
					Times(1).
					Return(nil, nil)
				querier.EXPECT().
					FindGroupByName(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				require.NotNil(t, recorder.Body)
			},
		},
		{
			name: "NotExistingGroup",
			url:  fmt.Sprintf("/blogs?gn=%s&key=%s", "not_existing_group", key),
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetAllBlogs(gomock.Any(), gomock.Any()).
					Times(0)
				querier.EXPECT().
					FindApiKeyByName(gomock.Any(), key).
					Times(1).
					Return(nil, nil)
				querier.EXPECT().
					FindGroupByName(gomock.Any(), gomock.Any()).
					Times(1).
					Return(&models.Group{}, fmt.Errorf("Failed to FindGroupByName: %w", sql.ErrNoRows))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name: "InternalDBErrorWhenReadingGroup",
			url:  fmt.Sprintf("/blogs?gn=%s&key=%s", groupName, key),
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetAllBlogs(gomock.Any(), groupName).
					Times(0)
				querier.EXPECT().
					FindApiKeyByName(gomock.Any(), key).
					Times(1).
					Return(nil, nil)
				querier.EXPECT().
					FindGroupByName(gomock.Any(), groupName).
					Times(1).
					Return(&models.Group{}, fmt.Errorf("Failed to FindGroupByName: %w", sql.ErrConnDone))
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name: "InternalDBError",
			url:  fmt.Sprintf("/blogs?gn=%s&key=%s", groupName, key),
			buildStubs: func(querier *mockdb.MockQuerier) {
				querier.EXPECT().
					GetAllBlogs(gomock.Any(), groupName).
					Times(1).
					Return(nil, errors.New("internal server error"))
				querier.EXPECT().
					FindApiKeyByName(gomock.Any(), key).
					Times(1).
					Return(nil, nil)
				querier.EXPECT().
					FindGroupByName(gomock.Any(), groupName).
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
