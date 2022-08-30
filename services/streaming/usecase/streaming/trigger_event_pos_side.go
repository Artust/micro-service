package streaming

import (
	"avatar/services/streaming/domain/broker"
	pb "avatar/services/streaming/protos"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"google.golang.org/grpc/metadata"
)

type CheckInResponse struct {
	Event   string
	Payload string
}

func ListenEventPOSSide(input *pb.ListenEventPOSSideRequest, stream pb.Streaming_ListenEventPOSSideServer, broker broker.Broker) error {
	ctx := stream.Context()
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return errors.New("missing metadata")
	}
	if len(md.Get("talkSessionId")) == 0 {
		return errors.New("missing talkSessionId")
	}
	talkSessionId := md.Get("talkSessionId")[0]
	dataChannel, errChannel, closeConsumer := broker.Consume(ctx, fmt.Sprintf("ListenEventPOSSide-%s", talkSessionId), true)
	for {
		select {
		case <-ctx.Done():
			err := closeConsumer()
			if err != nil {
				fmt.Println(err)
			}
			err = broker.DeleteTopic(fmt.Sprintf("ListenEventPOSSide-%s", talkSessionId))
			if err != nil {
				fmt.Println(err)
			}
			return ctx.Err()
		case data := <-dataChannel:
			var checkInResponse CheckInResponse
			err := json.Unmarshal(data, &checkInResponse)
			if err != nil {
				log.Println("error when unmarshal events:", err)
				return err
			}
			err = stream.Send(&pb.ListenEventPOSSideResponse{
				Event:   checkInResponse.Event,
				Payload: checkInResponse.Payload,
			})
			if err != nil {
				return err
			}
		case err := <-errChannel:
			return err
		}
	}
}
