package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	pb "avatar/services/gateway/protos/gateway"
)

func main() {
	streamingConnection, err := grpc.Dial("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
		return
	}
	streamingClient := pb.NewAvatarClient(streamingConnection)
	ctx := metadata.AppendToOutgoingContext(context.Background(), "talkSessionId", "1")
	stream, err := streamingClient.StreamPOSSideVoice(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := make([]byte, 16*1024)
	for {
		n, err := os.Stdin.Read(buf)
		if n > 0 {
			if err := stream.Send(&pb.Data{Data: buf[:n]}); err != nil {
				log.Printf("Could not send audio: %v", err)
			}
		}
		if err == io.EOF {
			// Nothing else to pipe, close the stream.
			if err := stream.CloseSend(); err != nil {
				log.Fatalf("Could not close stream: %v", err)
			}
			return
		}
		if err != nil {
			log.Printf("Could not read from stdin: %v", err)
			continue
		}
	}
}
