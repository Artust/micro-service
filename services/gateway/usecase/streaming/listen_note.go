package streaming

import (
	pb "avatar/services/gateway/protos/gateway"
	pbStreaming "avatar/services/gateway/protos/streaming"
	pbTalkSession "avatar/services/gateway/protos/talk_session"
	"avatar/services/streaming/domain/broker/event"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

type NotePayload struct {
	Id            int64  `json:"id"`
	TalkSessionId int64  `json:"talkSessionId"`
	Content       string `json:"content"`
	IsGuest       *bool  `json:"isGuest"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt"`
}

func ListenNotes(
	request *pb.Empty,
	stream pb.Avatar_ListenNotesServer,
	streamingClient pbStreaming.StreamingClient,
	talkSessionClient pbTalkSession.TalkSessionClient,
) error {
	ctx := stream.Context()
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return errors.New("missing metadata")
	}
	defaultTalkSessionId := 1
	var err error
	if len(md.Get("talkSessionId")) != 0 {
		defaultTalkSessionId, err = strconv.Atoi(md.Get("talkSessionId")[0])
		if err != nil {
			return err
		}
	}
	ctx = metadata.AppendToOutgoingContext(ctx, "talkSessionId", fmt.Sprintf("%d", defaultTalkSessionId))
	streamInternal, err := streamingClient.ListenNotes(ctx, &emptypb.Empty{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	data := &pbTalkSession.GetListNoteRequest{
		TalkSessionId: int64(defaultTalkSessionId),
		Page:          1,
		PerPage:       10000,
	}
	noteResponse, err := talkSessionClient.GetListNote(context.Background(), data)
	if err != nil {
		return err
	}
	var notePayload NotePayload
	for i := len(noteResponse.Notes) - 1; i >= 0; i-- {
		err = copier.Copy(&notePayload, noteResponse.Notes[i])
		if err != nil {
			return err
		}
		payload, err := json.Marshal(notePayload)
		if err != nil {
			return err
		}
		err = stream.Send(&pb.ListenEventResponse{
			Event:   int32(event.NoteEventCreate),
			Payload: string(payload),
		})
		if err != nil {
			fmt.Println(err)
			return err
		}
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
				err := stream.Send(&pb.ListenEventResponse{
					Event:   int32(data.Event),
					Payload: data.Payload,
				})
				if err != nil {
					fmt.Println(err)
					return err
				}
			}
		}
	}
}
