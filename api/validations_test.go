package api

import (
	"context"
	"database/sql"
	"errors"
	"os"
	"testing"

	mockdb "github.com/android-project-46group/api-server/db/mock"
	"github.com/android-project-46group/api-server/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestIsApiKeyValid(t *testing.T) {
	key := "valid_api_key"

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
					FindApiKeyByName(gomock.Any(), key).
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
					FindApiKeyByName(gomock.Any(), "invalid_key").
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
					FindApiKeyByName(gomock.Any(), gomock.Any()).
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
			logger, _, _ := util.NewStandardLogger("go-test", "api-saka", os.Stdout)

			server, err := NewServer(config, querier, matcher, logger, &mockGrpcClient{})
			require.NoError(t, err)

			result := server.isApiKeyValid(context.Background(), tc.key)
			require.Equal(t, tc.expected, result)
		})
	}
}
