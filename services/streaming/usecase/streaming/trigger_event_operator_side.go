package streaming

import (
	"avatar/services/streaming/domain/broker"
	"avatar/services/streaming/domain/entity"
	pb "avatar/services/streaming/protos"
	"encoding/json"
	"errors"
	"fmt"

	"google.golang.org/grpc/metadata"
)

func ListenEventOperatorSide(input *pb.ListenEventOperatorSideRequest, stream pb.Streaming_ListenEventOperatorSideServer, broker broker.Broker) error {
	ctx := stream.Context()
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return errors.New("missing metadata")
	}
	if len(md.Get("talkSessionId")) == 0 {
		return errors.New("missing talkSessionId")
	}
	talkSessionId := md.Get("talkSessionId")[0]
	dataChannel, errChannel, closeConsumer := broker.Consume(ctx, fmt.Sprintf("ListenEventOperatorSide-%s", talkSessionId), true)
	for {
		select {
		case <-ctx.Done():
			err := closeConsumer()
			if err != nil {
				fmt.Println(err)
			}
			err = broker.DeleteTopic(fmt.Sprintf("ListenEventOperatorSide-%s", talkSessionId))
			if err != nil {
				fmt.Println(err)
			}
			return ctx.Err()
		case data := <-dataChannel:
			var streamOperatorCheckIn entity.StreamOperatorCheckIn
			err := json.Unmarshal(data, &streamOperatorCheckIn)
			if err != nil {
				return err
			}
			err = stream.Send(&pb.ListenEventOperatorSideResponse{
				Event:   streamOperatorCheckIn.Event,
				Payload: streamOperatorCheckIn.Payload,
			})
			if err != nil {
				return err
			}
		case err := <-errChannel:
			return err
		}
	}
}
