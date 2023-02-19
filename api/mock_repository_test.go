package api

import (
	"context"
	"io"
)

type mockGrpcClient struct {
	HealthFunc          func(ctx context.Context) (string, error)
	DownloadMembersFunc func(ctx context.Context, writer io.Writer) error
}

func (m *mockGrpcClient) Health(ctx context.Context) (string, error) {
	return m.HealthFunc(ctx)
}

func (m *mockGrpcClient) DownloadMembers(ctx context.Context, writer io.Writer) error {
	return m.DownloadMembersFunc(ctx, writer)
}
