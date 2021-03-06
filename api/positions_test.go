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
	"github.com/android-project-46group/api-server/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetPositionsAPI(t *testing.T) {

	title := "nogizaka"
	key := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"

	testCases := []struct {
		name          string
		url           string
		buildStubs    func(store *mockdb.MockQuerier)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			url:  fmt.Sprintf("/positions?title=%s&key=%s", title, key),
			buildStubs: func(store *mockdb.MockQuerier) {
				store.EXPECT().
					GetPositionFromTitle(title).
					Times(1).
					Return([]db.PositionMemberBind{}, nil)
				store.EXPECT().
					FindApiKeyByName(key).
					Times(1).
					Return(nil, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name: "InvalidAPIKey",
			url:  fmt.Sprintf("/positions?title=%s&key=%s", title, "invalid_key"),
			buildStubs: func(store *mockdb.MockQuerier) {
				store.EXPECT().
					GetPositionFromTitle(gomock.Any()).
					Times(0)
				store.EXPECT().
					FindApiKeyByName("invalid_key").
					Times(1).
					Return(nil, sql.ErrNoRows)
				store.EXPECT().
					ExistGroup(gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusForbidden, recorder.Code)
			},
		},
		{
			name: "InternalDBError",
			url:  fmt.Sprintf("/positions?title=%s&key=%s", title, key),
			buildStubs: func(store *mockdb.MockQuerier) {
				store.EXPECT().
					GetPositionFromTitle(title).
					Times(1).
					Return(nil, errors.New("internal server error"))
				store.EXPECT().
					FindApiKeyByName(key).
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
			server, err := NewServer(config, querier)
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
