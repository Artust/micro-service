package handler

import (
	"avatar/services/streaming/config"
	"avatar/services/streaming/domain/broker"

	pb "avatar/services/streaming/protos"

	googleSpeech "cloud.google.com/go/speech/apiv1"
)

type StreamingServer struct {
	BrokerClient       broker.Broker
	GoogleSpeechClient *googleSpeech.Client
	Config             *config.Environment
	pb.UnimplementedStreamingServer
}

func CreateServer(
	brokerClient broker.Broker,
	googleSpeechClient *googleSpeech.Client,
	config *config.Environment,
) *StreamingServer {
	return &StreamingServer{
		BrokerClient:       brokerClient,
		GoogleSpeechClient: googleSpeechClient,
		Config:             config,
	}
}
