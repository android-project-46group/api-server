package grpc

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/android-project-46group/protobuf/gen/go/protobuf"
	"google.golang.org/grpc"
)

func (c *grpcClient) DownloadMembers(ctx context.Context, writer io.Writer) error {
	conn, err := grpc.Dial("localhost:9876", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewDownloadClient(conn)

	req := &pb.DownloadMembersZipRequest{}
	stream, err := client.DownloadMembersZip(ctx, req)
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
