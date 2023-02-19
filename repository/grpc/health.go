package grpc

import (
	"context"
	"fmt"

	pb "github.com/android-project-46group/protobuf/gen/go/protobuf"
)

func (g *grpcClient) Health(ctx context.Context) (string, error) {
	req := &pb.HealthRequest{}

	resp, err := g.downloadClient.Health(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to check health: %w", err)
	}

	return resp.GetMessage(), nil
}
