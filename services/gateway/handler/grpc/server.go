package grpc

import (
	pb "avatar/services/gateway/protos/gateway"
	pbStreaming "avatar/services/gateway/protos/streaming"
	pbTalkSession "avatar/services/gateway/protos/talk_session"
)

type GatewayServer struct {
	pb.UnimplementedAvatarServer
	StreamingClient   pbStreaming.StreamingClient
	TalkSessionClient pbTalkSession.TalkSessionClient
}

func CreateServer(
	streamingClient pbStreaming.StreamingClient,
	talkSessionClient pbTalkSession.TalkSessionClient,
) *GatewayServer {
	return &GatewayServer{
		StreamingClient:   streamingClient,
		TalkSessionClient: talkSessionClient,
	}
}
