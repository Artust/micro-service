package streaming

import (
	"avatar/services/streaming/domain/broker"
	pb "avatar/services/streaming/protos"
	"avatar/services/streaming/usecase/speech_to_text"
	"encoding/json"
	"errors"
	"fmt"

	"google.golang.org/grpc/metadata"
)

func SpeechToText(
	stream pb.Streaming_SpeechToTextServer,
	broker broker.Broker,
) error {
	ctx := stream.Context()
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return errors.New("missing metadata")
	}
	if len(md.Get("talkSessionId")) == 0 {
		return errors.New("missing talkSessionId")
	}
	talkSessionId := md.Get("talkSessionId")[0]
	dataChannel, errChannel, closeConsumer := broker.Consume(ctx, fmt.Sprintf("SpeechToText-%v", talkSessionId), true)
	for {
		select {
		case <-ctx.Done():
			err := closeConsumer()
			if err != nil {
				fmt.Println(err)
			}
			err = broker.DeleteTopic(fmt.Sprintf("POSToOp-SpeechToText-%v", talkSessionId))
			if err != nil {
				fmt.Println(err)
			}
			return ctx.Err()
		case rawData := <-dataChannel:
			var speechToTextMessage speech_to_text.SpeechToTextMessage
			err := json.Unmarshal(rawData, &speechToTextMessage)
			if err != nil {
				fmt.Println(err)
			} else {
				dataSend := &pb.SpeechToTextData{
					Speaker:     string(speechToTextMessage.Speaker),
					Content:     speechToTextMessage.Content,
					SendingTime: speechToTextMessage.SendingTime,
				}
				err := stream.Send(dataSend)
				if err != nil {
					return err
				}
			}
		case err := <-errChannel:
			return err
		}
	}
}
