package grpc

import (
	"context"
	"fmt"
	"io"

	pb "github.com/android-project-46group/protobuf/gen/go/protobuf"
)

func (c *grpcClient) DownloadMembers(ctx context.Context, writer io.Writer) error {
	req := &pb.DownloadMembersZipRequest{}

	stream, err := c.downloadClient.DownloadMembersZip(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to DownloadMembersZip: %w", err)
	}

	for {
		feature, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to stream.Recv: %w", err)
		}

		writer.Write(feature.GetData())
	}

	return nil
}
