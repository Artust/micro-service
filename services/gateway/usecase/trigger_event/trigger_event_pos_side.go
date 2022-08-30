package trigger_event

import (
	"avatar/services/gateway/config"
	"avatar/services/gateway/domain/broker"
	pb "avatar/services/gateway/protos/pos"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

type TriggerEventPOSSideInput struct {
	Event   string `json:"event" binding:"required"`
	Payload string `json:"payload" binding:"required"`
}
type PayloadTriggerEventPOSRequest struct {
	TalkSessionId int64 `json:"talkSessionId"`
}

type TrigerEventGuestCheckInPayload struct {
	ID            int64
	AnimationFile string
	SoundFile     string
}

func TriggerEventPOSSide(
	input *TriggerEventPOSSideInput,
	posClient pb.POSClient,
	config *config.Environment,
	broker broker.Broker) error {
	var payloadInput PayloadTriggerEventPOSRequest
	err := json.Unmarshal([]byte(input.Payload), &payloadInput)
	if err != nil {
		fmt.Println("error:", err)
	}
	routine, err := posClient.GetListRoutine(context.Background(), &pb.GetListRoutineRequest{
		// PosId:      input.PosId,
		// CategoryId: input.CategoryId,
		Page:    1,
		PerPage: 1,
	})
	if err != nil {
		return err
	}
	payload := &TrigerEventGuestCheckInPayload{
		ID:            routine.GetListRoutineResponse[0].Id,
		AnimationFile: fmt.Sprintf("%s%s", config.S3Uri, routine.GetListRoutineResponse[0].AnimationFile),
		SoundFile:     fmt.Sprintf("%s%s", config.S3Uri, routine.GetListRoutineResponse[0].SoundFile),
	}
	b, err := json.Marshal(payload)
	if err != nil {
		log.Println("error when mashal routine:", err)
		return err
	}
	response := &TriggerEventPOSSideInput{
		Event:   input.Event,
		Payload: string(b),
	}
	b, err = json.Marshal(response)
	if err != nil {
		log.Println("error when mashal payload routine:", err)
		return err
	}

	err = broker.Produce(fmt.Sprintf("ListenEventPOSSide-%d", payloadInput.TalkSessionId), b)
	if err != nil {
		log.Println("error produce ListenEventPOSSide:", err)
		return err
	}
	producerOperator := &TriggerEventPOSSideInput{
		Event:   "GuestCheckIn",
		Payload: fmt.Sprintf(`{"posId":%v}`, 1),
	}
	b, err = json.Marshal(producerOperator)
	if err != nil {
		log.Println("error when mashal payload routine:", err)
		return err
	}

	err = broker.Produce(fmt.Sprintf("ListenEventOperatorSide-%d", payloadInput.TalkSessionId), b)
	if err != nil {
		log.Println("error produce listenEventOperatorSide:", err)
		return err
	}
	return nil
}
