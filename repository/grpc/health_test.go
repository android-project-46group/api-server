package grpc_test

import (
	"context"
	"os"
	"testing"

	"github.com/android-project-46group/api-server/repository/grpc"
	"github.com/android-project-46group/api-server/util"
	pb "github.com/android-project-46group/protobuf/gen/go/protobuf"
	mockgrpc "github.com/android-project-46group/protobuf/gen/go/protobuf/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHealthGrpc(t *testing.T) {

	testCases := []struct {
		name          string
		buildStubs    func(mock *mockgrpc.MockDownloadClient)
		checkResponse func(t *testing.T, response string, err error)
	}{
		{
			name: "OK",
			buildStubs: func(mock *mockgrpc.MockDownloadClient) {
				mock.EXPECT().
					Health(gomock.Any(), gomock.Any()).
					Times(1).
					Return(&pb.HealthReply{Message: "status ok"}, nil)
			},
			checkResponse: func(t *testing.T, response string, err error) {
				assert.Nil(t, err)
				assert.Equal(t, "status ok", response)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			config := util.Config{
				Service: "saka-api",
				GRPCURL: "localhost:12345",
			}
			gmock := mockgrpc.NewMockDownloadClient(ctrl)

			tc.buildStubs(gmock)

			logger, _, _ := util.NewStandardLogger("go-test", "api-saka", os.Stdout)
			gclient := grpc.NewClient(logger, config, gmock)

			// Act
			response, err := gclient.Health(context.Background())

			// Assert
			tc.checkResponse(t, response, err)
		})
	}
}
