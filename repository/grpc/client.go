package grpc

import (
	"context"
	"io"

	"github.com/android-project-46group/api-server/util"
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
	logger  util.Logger
	grpcURL string
}

func NewClient(logger util.Logger, config util.Config) GrpcClient {
	return &grpcClient{
		logger:  logger,
		grpcURL: config.GRPCURL,
	}
}
