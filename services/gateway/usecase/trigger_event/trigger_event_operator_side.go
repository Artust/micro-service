package trigger_event

import (
	"avatar/services/gateway/domain/broker"
	pb "avatar/services/gateway/protos/pos"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

type TriggerEventOperatorSideInput struct {
	Event   string `json:"event" binding:"required"`
	Payload string `json:"payload" binding:"required"`
}
type TriggerEventOperatorSideOutPut struct {
	Event   string `json:"event"`
	Payload string `json:"payload"`
}
type PayloadTriggerEventOperatorRequest struct {
	TalkSessionId int64 `json:"talkSessionId"`
}

func TriggerEventOperatorSide(input *TriggerEventOperatorSideInput,
	posClient pb.POSClient,
	broker broker.Broker) (*TriggerEventOperatorSideOutPut, error) {
	var payloadInput PayloadTriggerEventOperatorRequest
	err := json.Unmarshal([]byte(input.Payload), &payloadInput)
	if err != nil {
		fmt.Println("error:", err)
	}
	routine, err := posClient.TriggerEventOperatorSide(context.Background(), &pb.TriggerEventOperatorSideRequest{
		Event:   input.Event,
		Payload: input.Payload,
	})
	if err != nil {
		return nil, err
	}
	response := &TriggerEventOperatorSideOutPut{
		Event:   routine.Event,
		Payload: routine.Payload,
	}
	b, err := json.Marshal(response)
	if err != nil {
		log.Println("error when mashal payload routine:", err)
		return nil, err
	}
	err = broker.Produce(fmt.Sprintf("ListenEventPOSSide-%d", payloadInput.TalkSessionId), b)
	if err != nil {
		log.Println("error produce ListenEventPOSSide:", err)
		return nil, err
	}
	return response, nil
}
