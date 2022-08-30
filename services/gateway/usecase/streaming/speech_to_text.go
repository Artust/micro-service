package streaming

import (
	pb "avatar/services/gateway/protos/gateway"
	pbStreaming "avatar/services/gateway/protos/streaming"
	"errors"
	"fmt"
	"io"
	"time"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SpeechToTextMessage struct {
	Speaker     string `json:"speaker"`
	Content     string `json:"content"`
	SendingTime string `json:"sendingTime"`
}

func SpeechToText(
	stream pb.Avatar_SpeechToTextServer,
	streamingClient pbStreaming.StreamingClient,
) error {
	ctx := stream.Context()
	request := &emptypb.Empty{}
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return errors.New("missing metadata")
	}
	if len(md.Get("talkSessionId")) == 0 {
		ctx = metadata.AppendToOutgoingContext(ctx, "talkSessionId", "1")
	} else {
		ctx = metadata.AppendToOutgoingContext(ctx, "talkSessionId", md.Get("talkSessionId")[0])
	}
	streamInternal, err := streamingClient.SpeechToText(ctx, request)
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
			if err != nil && err != io.EOF {
				fmt.Println(err)
				return err
			}
			if err == io.EOF {
				time.Sleep(500 * time.Millisecond)
			}
			if data != nil {
				err := stream.Send(&pb.SpeechToTextData{
					Speaker:     data.Speaker,
					Content:     data.Content,
					SendingTime: data.SendingTime,
				})
				if err != nil {
					fmt.Println(err)
					return err
				}
			}
		}
	}
}
