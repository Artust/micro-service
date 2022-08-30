package streaming

import (
	pb "avatar/services/gateway/protos/gateway"
	pbStreaming "avatar/services/gateway/protos/streaming"
	"errors"
	"fmt"
	"io"
	"time"

	"google.golang.org/grpc/metadata"
)

func ListenEventOperatorSide(request *pb.Empty, stream pb.Avatar_ListenEventOperatorSideServer, streamingClient pbStreaming.StreamingClient) error {
	ctx := stream.Context()
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return errors.New("missing metadata")
	}
	if len(md.Get("talkSessionId")) == 0 {
		ctx = metadata.AppendToOutgoingContext(ctx, "talkSessionId", "1")
	} else {
		ctx = metadata.AppendToOutgoingContext(ctx, "talkSessionId", md.Get("talkSessionId")[0])
	}
	streamInternal, err := streamingClient.ListenEventOperatorSide(ctx, &pbStreaming.ListenEventOperatorSideRequest{
		CenterID: "1",
	})
	if err != nil {
		fmt.Println(err)
		return err
	}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			data, err := streamInternal.Recv()
			if err == io.EOF {
				time.Sleep(500 * time.Millisecond)
			}
			if data != nil && len(data.Event) > 0 {
				err := stream.Send(&pb.ListenEventOperatorSideResponse{
					Event:   data.Event,
					Payload: data.Payload,
				})
				if err != nil {
					fmt.Println(err)
					return err
				}
			}
		}
	}
}
