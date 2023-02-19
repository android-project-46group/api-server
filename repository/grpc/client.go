package grpc

import (
	"context"
	"fmt"
	"io"

	"github.com/android-project-46group/api-server/util"
	pb "github.com/android-project-46group/protobuf/gen/go/protobuf"
	"google.golang.org/grpc"
)

type GrpcClient interface {
	// ヘルスチェック用のエンドポイント。
	Health(ctx context.Context) (string, error)

	// メンバー情報をダウンロードするエンドポイント。
	//
	// zip 化されたメンバー情報を byte 配列で取得する。
	DownloadMembers(ctx context.Context, writer io.Writer) error
}

type grpcClient struct {
	logger         util.Logger
	grpcURL        string
	downloadClient pb.DownloadClient
}

func NewClient(logger util.Logger, config util.Config, downloadClient pb.DownloadClient) GrpcClient {

	return &grpcClient{
		logger:         logger,
		grpcURL:        config.GRPCURL,
		downloadClient: downloadClient,
	}
}

func NewGRPCClient(grpcURL string) (pb.DownloadClient, func(), error) {
	conn, err := grpc.Dial(grpcURL, grpc.WithInsecure())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to dial: %w", err)
	}

	return pb.NewDownloadClient(conn), func() {
		conn.Close()
	}, nil
}
