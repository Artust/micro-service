package streaming

import (
	"avatar/services/streaming/domain/broker"
	"avatar/services/streaming/domain/broker/topic"
	pb "avatar/services/streaming/protos"
	"encoding/json"
	"errors"
	"fmt"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

func ListenNotes(request *emptypb.Empty, stream pb.Streaming_ListenNotesServer, broker broker.Broker) error {
	ctx := stream.Context()
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return errors.New("missing metadata")
	}
	if len(md.Get("talkSessionId")) == 0 {
		return errors.New("missing talkSessionId")
	}
	talkSessionId := md.Get("talkSessionId")[0]
	topic := fmt.Sprintf("%s-%s", topic.ListenNoteOperatorSideTopic, talkSessionId)
	dataChannel, errChannel, closeConsumer := broker.Consume(ctx, topic, true)
	for {
		select {
		case <-ctx.Done():
			err := closeConsumer()
			if err != nil {
				fmt.Println(err)
			}
			err = broker.DeleteTopic(topic)
			if err != nil {
				fmt.Println(err)
			}
			return ctx.Err()
		case data := <-dataChannel:
			var listenEventResponse pb.ListenEventResponse
			err := json.Unmarshal(data, &listenEventResponse)
			if err != nil {
				return err
			}
			err = stream.Send(&listenEventResponse)
			if err != nil {
				return err
			}
		case err := <-errChannel:
			return err
		}
	}
}
