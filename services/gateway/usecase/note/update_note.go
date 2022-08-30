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

type UpdateNoteInput struct {
	Id      int64
	Content string `json:"content"`
	IsGuest bool   `json:"isGuest"`
}

func UpdateNote(
	input *UpdateNoteInput,
	talkSessionClient pb.TalkSessionClient,
	broker broker.Broker,
) (*CreateNoteOutput, error) {
	data := &pb.UpdateNoteRequest{
		Id:      input.Id,
		Content: input.Content,
		IsGuest: input.IsGuest,
	}
	response, err := talkSessionClient.UpdateNote(context.Background(), data)
	if err != nil {
		return nil, err
	}
	note := &CreateNoteOutput{
		Id:            response.Id,
		TalkSessionId: response.TalkSessionId,
		Content:       response.Content,
		IsGuest:       &response.IsGuest,
		OperatorId:    10,
		OperatorName:  "Tran Manh Tien",
		CreatedAt:     response.CreatedAt,
		UpdatedAt:     response.UpdatedAt,
	}
	payload, err := json.Marshal(note)
	if err != nil {
		log.Error("error mashal:", err)
		return nil, err
	}
	message, err := json.Marshal(NoteEvent{
		Event:   int(event.NoteEventUpdate),
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

	return note, nil
}
