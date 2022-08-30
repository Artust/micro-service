package handler

import (
	pb "avatar/services/streaming/protos"
	"avatar/services/streaming/usecase/streaming"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *StreamingServer) StreamOperatorSideVideo(stream pb.Streaming_StreamOperatorSideVideoServer) error {
	return streaming.StreamOperatorSideVideo(stream, server.BrokerClient)
}

func (server *StreamingServer) StreamOperatorSideVoice(stream pb.Streaming_StreamOperatorSideVoiceServer) error {
	return streaming.StreamOperatorSideVoice(stream, server.BrokerClient, server.Config, server.GoogleSpeechClient)
}

func (server *StreamingServer) StreamPOSSideVideo(stream pb.Streaming_StreamPOSSideVideoServer) error {
	return streaming.StreamPOSSideVideo(stream, server.BrokerClient)
}

func (server *StreamingServer) StreamPOSSideVoice(stream pb.Streaming_StreamPOSSideVoiceServer) error {
	return streaming.StreamPOSSideVoice(stream, server.BrokerClient, server.Config, server.GoogleSpeechClient)
}

func (server *StreamingServer) SpeechToText(request *emptypb.Empty, stream pb.Streaming_SpeechToTextServer) error {
	return streaming.SpeechToText(stream, server.BrokerClient)
}

func (server *StreamingServer) ListenEventPOSSide(request *pb.ListenEventPOSSideRequest, stream pb.Streaming_ListenEventPOSSideServer) error {
	return streaming.ListenEventPOSSide(request, stream, server.BrokerClient)
}

func (server *StreamingServer) ListenEventOperatorSide(request *pb.ListenEventOperatorSideRequest, stream pb.Streaming_ListenEventOperatorSideServer) error {
	return streaming.ListenEventOperatorSide(request, stream, server.BrokerClient)
}

func (server *StreamingServer) ListenNotes(request *emptypb.Empty, stream pb.Streaming_ListenNotesServer) error {
	return streaming.ListenNotes(request, stream, server.BrokerClient)
}

func (server *StreamingServer) ListenListPos(request *pb.ListenListPosRequest, stream pb.Streaming_ListenListPosServer) error {
	return streaming.ListenListPos(request, stream, server.BrokerClient)
}
