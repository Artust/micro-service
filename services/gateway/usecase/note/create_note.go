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

type NoteEvent struct {
	Event   int    `json:"event"`
	Payload string `json:"payload"`
}

type CreateNoteInput struct {
	TalkSessionId int64  `json:"talkSessionId" binding:"required"`
	Content       string `json:"content" binding:"required"`
	IsGuest       bool   `json:"isGuest"`
}

type CreateNoteOutput struct {
	Id            int64  `json:"id"`
	TalkSessionId int64  `json:"talkSessionId"`
	Content       string `json:"content"`
	IsGuest       *bool  `json:"isGuest"`
	OperatorId    int64  `json:"operatorId"`
	OperatorName  string `json:"operatorName"`
	CreatedAt     string `json:"createdAt"`
	UpdatedAt     string `json:"updatedAt,omitempty"`
}

func CreateNote(
	input *CreateNoteInput,
	talkSessionClient pb.TalkSessionClient,
	broker broker.Broker,
) (*CreateNoteOutput, error) {
	data := &pb.CreateNoteRequest{
		TalkSessionId: input.TalkSessionId,
		Content:       input.Content,
		IsGuest:       input.IsGuest,
	}
	response, err := talkSessionClient.CreateNote(context.Background(), data)
	if err != nil {
		return nil, err
	}
	output := &CreateNoteOutput{
		Id:            response.Id,
		TalkSessionId: response.TalkSessionId,
		Content:       response.Content,
		IsGuest:       &response.IsGuest,
		OperatorId:    10,
		OperatorName:  "Tran Manh Tien",
		CreatedAt:     response.CreatedAt,
		UpdatedAt:     response.UpdatedAt,
	}
	payload, err := json.Marshal(output)
	if err != nil {
		log.Error("error mashal:", err)
		return nil, err
	}
	message, err := json.Marshal(NoteEvent{
		Event:   int(event.NoteEventCreate),
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
	return output, nil
}
