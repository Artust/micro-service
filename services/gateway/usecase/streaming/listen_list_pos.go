package streaming

import (
	pb "avatar/services/gateway/protos/gateway"
	pbStreaming "avatar/services/gateway/protos/streaming"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

func ListenListPos(request *pb.Empty, stream pb.Avatar_ListenListPosServer, streamingClient pbStreaming.StreamingClient) error {
	ctx := stream.Context()
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return errors.New("missing metadata")
	}
	token := md.Get("token")
	log.Println(token)
	streamInternal, err := streamingClient.ListenListPos(ctx, &pbStreaming.ListenListPosRequest{
		GroupId: 0,
	})
	if err != nil {
		log.Error(err)
		return err
	}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			data, err := streamInternal.Recv()
			if err == io.EOF {
				time.Sleep(500 * time.Millisecond)
			}
			if data != nil {
				response := &pb.ListenListPosResponse{}
				var listPosResponse pb.ListPosResponse
				for _, val := range data.ListPosResponse {
					err := copier.Copy(&val, &listPosResponse)
					if err != nil {
						return err
					}
					response.ListPosResponse = append(response.ListPosResponse, &listPosResponse)
				}
				err = stream.Send(response)
				if err != nil {
					fmt.Println(err)
					return err
				}
			}
		}
	}
}
