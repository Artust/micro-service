package note

import (
	"avatar/services/gateway/domain/broker"
	"avatar/services/gateway/domain/broker/event"
	"avatar/services/gateway/domain/broker/topic"
	pb "avatar/services/gateway/protos/talk_session"
	"context"
	"encoding/json"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type DeleteNoteInput struct {
	Id int64 `json:"id" binding:"required"`
}

type DeleteNoteOutput struct {
	RowsAffected  int64 `json:"rowsAffected"`
	TalkSessionId int64 `json:"talkSessionId,omitempty"`
}

func DeleteNote(
	input *DeleteNoteInput,
	talkSessionClient pb.TalkSessionClient,
	broker broker.Broker,
) (*DeleteNoteOutput, error) {
	data := &pb.DeleteNoteRequest{
		Id: input.Id,
	}
	response, err := talkSessionClient.DeleteNote(context.Background(), data)
	if err != nil {
		return nil, err
	}
	payload, err := json.Marshal(data)
	if err != nil {
		log.Error("error mashal:", err)
		return nil, err
	}
	message, err := json.Marshal(NoteEvent{
		Event:   int(event.NoteEventDelete),
		Payload: string(payload),
	})
	if err != nil {
		log.Error("error mashal:", err)
		return nil, err
	}
	topic := fmt.Sprintf("%s-%d", topic.ListenNoteOperatorSideTopic, response.TalkSessionId)
	err = broker.Produce(topic, message)
	if err != nil {
		log.Errorf("error produce %s:%v", topic, err)
		return nil, err
	}
	output := &DeleteNoteOutput{
		RowsAffected: response.RowsAffected,
	}
	return output, nil
}
