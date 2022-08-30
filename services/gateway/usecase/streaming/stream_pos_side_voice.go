package streaming

import (
	pb "avatar/services/gateway/protos/gateway"
	pbStreaming "avatar/services/gateway/protos/streaming"
	"errors"
	"fmt"
	"io"
	"time"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/metadata"
)

func StreamPOSSideVoice(
	stream pb.Avatar_StreamPOSSideVoiceServer,
	streamingClient pbStreaming.StreamingClient,
) error {
	eg, ctx := errgroup.WithContext(stream.Context())
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return errors.New("missing metadata")
	}
	if len(md.Get("talkSessionId")) == 0 {
		ctx = metadata.AppendToOutgoingContext(ctx, "talkSessionId", "1")
	} else {
		ctx = metadata.AppendToOutgoingContext(ctx, "talkSessionId", md.Get("talkSessionId")[0])
	}
	streamInternal, err := streamingClient.StreamPOSSideVoice(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				data, err := streamInternal.Recv()
				if err != nil && err != io.EOF {
					return err
				}
				if err == io.EOF {
					time.Sleep(500 * time.Millisecond)
				}
				if data != nil && len(data.Data) > 0 {
					err := stream.Send(&pb.Data{Data: data.Data})
					if err != nil {
						fmt.Println(err)
						return err
					}
				}
			}
		}
	})
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				data, err := stream.Recv()
				if err != nil && err != io.EOF {
					return err
				}
				if err == io.EOF {
					time.Sleep(500 * time.Millisecond)
				}
				if data != nil && len(data.Data) > 0 {
					err := streamInternal.Send(&pbStreaming.Data{Data: data.Data})
					if err != nil {
						fmt.Println(err)
						return err
					}
				}
			}
		}
	})
	return eg.Wait()
}
