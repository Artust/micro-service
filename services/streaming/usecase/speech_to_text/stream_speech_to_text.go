package speech_to_text

import (
	"avatar/services/streaming/config"
	"avatar/services/streaming/domain/broker"
	"context"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"time"

	googleSpeech "cloud.google.com/go/speech/apiv1"
	"golang.org/x/sync/errgroup"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1"
)

func initStream(
	cfg *config.Environment,
	googleSpeechClient *googleSpeech.Client,
	ctx context.Context,
) (speechpb.Speech_StreamingRecognizeClient, error) {
	speechStream, err := googleSpeechClient.StreamingRecognize(ctx)
	if err != nil {
		return nil, err
	}
	err = speechStream.Send(&speechpb.StreamingRecognizeRequest{
		StreamingRequest: &speechpb.StreamingRecognizeRequest_StreamingConfig{
			StreamingConfig: &speechpb.StreamingRecognitionConfig{
				Config: &speechpb.RecognitionConfig{
					Encoding:        speechpb.RecognitionConfig_LINEAR16,
					SampleRateHertz: int32(cfg.SpeechToTextSampleRate),
					LanguageCode:    cfg.SpeechToTextLanguageCode,
					// AlternativeLanguageCodes: []string{
					// 	"vi-VN",
					// 	"ja-JP",
					// },
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return speechStream, nil
}

func StreamSpeechToText(
	cfg *config.Environment,
	googleSpeechClient *googleSpeech.Client,
	ctx context.Context,
	speechChannel chan []byte,
	broker broker.Broker,
	topic string,
	speaker Speaker,
) error {
	speechStream, err := initStream(cfg, googleSpeechClient, ctx)
	if err != nil {
		return err
	}
	eg, newCtx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		for {
			select {
			case <-newCtx.Done():
				return newCtx.Err()
			case rawSpeechData := <-speechChannel:
				speechData := convertF32leToS16leByte(rawSpeechData)
				err := speechStream.Send(&speechpb.StreamingRecognizeRequest{
					StreamingRequest: &speechpb.StreamingRecognizeRequest_AudioContent{
						AudioContent: speechData,
					},
				})
				if err != nil {
					return err
				}
			}
		}
	})
	eg.Go(func() error {
		for {
			select {
			case <-newCtx.Done():
				return newCtx.Err()
			default:
				data, err := speechStream.Recv()
				if err != nil && err != io.EOF {
					return err
				}
				if err := data.Error; err != nil {
					if err.Code == 11 {
						if err := speechStream.CloseSend(); err != nil {
							return err
						}
						newSpeechStream, err := initStream(cfg, googleSpeechClient, ctx)
						if err != nil {
							return err
						}
						speechStream = newSpeechStream
					} else {
						fmt.Printf("Could not recognize: %v", err)
						return errors.New(err.Message)
					}
				} else {
					for _, result := range data.Results {
						message := SpeechToTextMessage{
							Speaker:     speaker,
							Content:     result.Alternatives[0].Transcript,
							SendingTime: time.Now().Format(time.RFC3339),
						}
						jsonMessage, _ := json.Marshal(message)
						broker.Produce(topic, jsonMessage)
						fmt.Printf("Result: %v\n", result.Alternatives[0].Transcript)
					}
				}
			}
		}
	})
	return eg.Wait()
}

func convertF32leToS16leByte(data []byte) []byte {
	lengthData := len(data) / 2
	if lengthData%2 != 0 {
		lengthData += 1
	}
	newByte := make([]byte, lengthData)
	j := 0
	for i := 0; i < len(data); i += 4 {
		m := binary.LittleEndian.Uint32(data[i : i+4])
		m1 := math.Float32frombits(m)
		m2 := uint16(m1 * 32767)
		bytes := make([]byte, 2)
		binary.LittleEndian.PutUint16(bytes, m2)
		newByte[j] = bytes[0]
		newByte[j+1] = bytes[1]
		j += 2
	}
	return newByte
}
