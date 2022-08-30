package grpc

import (
	pb "avatar/services/gateway/protos/gateway"
	"avatar/services/gateway/usecase/streaming"
)

func (server *GatewayServer) StreamOperatorSideVideo(stream pb.Avatar_StreamOperatorSideVideoServer) error {
	return streaming.StreamOperatorSideVideo(stream, server.StreamingClient)
}

func (server *GatewayServer) StreamPOSSideVideo(stream pb.Avatar_StreamPOSSideVideoServer) error {
	return streaming.StreamPOSSideVideo(stream, server.StreamingClient)
}

func (server *GatewayServer) StreamOperatorSideVoice(stream pb.Avatar_StreamOperatorSideVoiceServer) error {
	return streaming.StreamOperatorSideVoice(stream, server.StreamingClient)
}

func (server *GatewayServer) StreamPOSSideVoice(stream pb.Avatar_StreamPOSSideVoiceServer) error {
	return streaming.StreamPOSSideVoice(stream, server.StreamingClient)
}

func (server *GatewayServer) SpeechToText(request *pb.Empty, stream pb.Avatar_SpeechToTextServer) error {
	return streaming.SpeechToText(stream, server.StreamingClient)
}

func (server *GatewayServer) ListenEventPOSSide(request *pb.Empty, stream pb.Avatar_ListenEventPOSSideServer) error {
	return streaming.ListenEventPOSSide(request, stream, server.StreamingClient)
}

func (server *GatewayServer) ListenEventOperatorSide(request *pb.Empty, stream pb.Avatar_ListenEventOperatorSideServer) error {
	return streaming.ListenEventOperatorSide(request, stream, server.StreamingClient)
}
func (server *GatewayServer) ListenNotes(request *pb.Empty, stream pb.Avatar_ListenNotesServer) error {
	return streaming.ListenNotes(request, stream, server.StreamingClient, server.TalkSessionClient)
}

func (server *GatewayServer) ListenListPos(request *pb.Empty, stream pb.Avatar_ListenListPosServer) error {
	return streaming.ListenListPos(request, stream, server.StreamingClient)
}
