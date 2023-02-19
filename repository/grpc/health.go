package grpc

import (
	"context"
	"fmt"

	pb "github.com/android-project-46group/protobuf/gen/go/protobuf"
	"google.golang.org/grpc"
)

func (g *grpcClient) Health(ctx context.Context) (string, error) {
	conn, err := grpc.Dial(g.grpcURL, grpc.WithInsecure())
	if err != nil {
		return "", fmt.Errorf("failed to dial: %w", err)
	}
	defer conn.Close()

	client := pb.NewDownloadClient(conn)

	req := &pb.HealthRequest{}
	resp, err := client.Health(ctx, req)
	if err != nil {
		return "", fmt.Errorf("failed to check health: %w", err)
	}

	return resp.GetMessage(), nil
}
