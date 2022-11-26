package api

import (
	"database/sql"
	"errors"
	"testing"

	mockdb "github.com/android-project-46group/api-server/db/mock"
	"github.com/android-project-46group/api-server/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestIsApiKeyValid(t *testing.T) {

	key := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"

	testCases := []struct {
		name       string
		key        string
		buildStubs func(store *mockdb.MockQuerier)
		expected   error
	}{
		{
			name: "OK",
			key:  key,
			buildStubs: func(store *mockdb.MockQuerier) {
				store.EXPECT().
					FindApiKeyByName(key).
					Times(1).
					Return(nil, nil)
			},
			expected: nil,
		},
		{
			name: "InvalidKey",
			key:  "invalid_key",
			buildStubs: func(store *mockdb.MockQuerier) {
				store.EXPECT().
					FindApiKeyByName("invalid_key").
					Times(1).
					Return(nil, sql.ErrNoRows)
			},
			expected: sql.ErrNoRows,
		},
		{
			name: "EmptyKey",
			key:  "",
			buildStubs: func(store *mockdb.MockQuerier) {
				store.EXPECT().
					FindApiKeyByName(gomock.Any()).
					Times(0)
			},
			expected: errors.New("API key is empty"),
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
			server, err := NewServer(config, querier, matcher)
			require.NoError(t, err)

			result := server.isApiKeyValid(tc.key)
			require.Equal(t, tc.expected, result)
		})
	}
}
