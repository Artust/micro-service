package trigger_event

import (
	"avatar/services/pos/config"
	"avatar/services/pos/domain/entity"
	"avatar/services/pos/domain/repository"
	pb "avatar/services/pos/protos"
	"context"
	"encoding/json"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func TriggerEventOperatorSide(
	ctx context.Context,
	db neo4j.Driver,
	routineRepository repository.RoutineRepository,
	input *pb.TriggerEventOperatorSideRequest,
	cfg *config.Environment,
) (*pb.TriggerEventOperatorSideResponse, error) {
	var triggerEventOperatorSidePayload entity.TriggerEventOperatorSidePayload
	err := json.Unmarshal([]byte(input.Payload), &triggerEventOperatorSidePayload)
	if err != nil {
		log.Error("error:", err)
	}
	session := db.NewSession(neo4j.SessionConfig{})
	routineRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		routine, err := routineRepository.GetById(ctx, triggerEventOperatorSidePayload.RoutineID)
		if err != nil {
			log.Error("error when get routine, error: ", err)
			return nil, err
		}
		return routine, nil
	})
	if err != nil {
		log.Error("error when get routine, error: ", err)
		return nil, err
	}
	routine := routineRaw.(*entity.PosRoutine)
	payload := &entity.TrigerEventPayload{
		ID:            routine.Id,
		AnimationFile: fmt.Sprintf("%s%s", cfg.S3URI, routine.AnimationFile),
		SoundFile:     fmt.Sprintf("%s%s", cfg.S3URI, routine.SoundFile),
	}
	b, err := json.Marshal(payload)
	if err != nil {
		log.Error("error when mashal routine:", err)
		return nil, err
	}
	response := &pb.TriggerEventOperatorSideResponse{
		Event:   input.Event,
		Payload: string(b),
	}

	if err != nil {
		fmt.Println("can not producer triger event kafka:", err)
		return nil, err
	}
	return response, nil
}
