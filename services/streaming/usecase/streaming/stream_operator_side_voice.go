package streaming

import (
	errUtil "avatar/pkg/err"
	"avatar/services/streaming/config"
	"avatar/services/streaming/domain/broker"
	pb "avatar/services/streaming/protos"
	"avatar/services/streaming/usecase/speech_to_text"
	"errors"
	"fmt"
	"io"
	"time"

	googleSpeech "cloud.google.com/go/speech/apiv1"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/metadata"
)

func StreamOperatorSideVoice(
	stream pb.Streaming_StreamOperatorSideVoiceServer,
	broker broker.Broker,
	cfg *config.Environment,
	googleSpeechClient *googleSpeech.Client,
) error {
	eg, ctx := errgroup.WithContext(stream.Context())
	speechChannel := make(chan []byte)
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return errors.New("missing metadata")
	}
	if len(md.Get("talkSessionId")) == 0 {
		return errors.New("missing talkSessionId")
	}
	talkSessionId := md.Get("talkSessionId")[0]
	eg.Go(errUtil.RecoverPanic(func() error {
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
					return nil
				}
				if data != nil && len(data.Data) > 0 {
					broker.Produce(fmt.Sprintf("OpToPOS-Voice-%v", talkSessionId), data.Data)
					speechChannel <- data.Data
				}
			}
		}
	}))
	eg.Go(errUtil.RecoverPanic(func() error {
		dataChannel, errChannel, closeConsumer := broker.Consume(ctx, fmt.Sprintf("POSToOp-Voice-%v", talkSessionId), true)
		for {
			select {
			case <-ctx.Done():
				err := closeConsumer()
				if err != nil {
					fmt.Println(err)
				}
				err = broker.DeleteTopic(fmt.Sprintf("POSToOp-Voice-%v", talkSessionId))
				if err != nil {
					fmt.Println(err)
				}
				return ctx.Err()
			case data := <-dataChannel:
				dataSend := &pb.Data{Data: data}
				err := stream.Send(dataSend)
				if err != nil {
					return err
				}
			case err := <-errChannel:
				return err
			}
		}
	}))
	eg.Go(errUtil.RecoverPanic(func() error {
		return speech_to_text.StreamSpeechToText(
			cfg,
			googleSpeechClient,
			ctx,
			speechChannel,
			broker,
			fmt.Sprintf("SpeechToText-%v", talkSessionId),
			speech_to_text.SpeakerOperator,
		)
	}))
	return eg.Wait()
}
