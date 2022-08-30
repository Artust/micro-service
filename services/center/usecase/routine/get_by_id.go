package routine

import (
	"avatar/services/center/config"
	"avatar/services/center/domain/entity"
	"avatar/services/center/domain/repository"
	pb "avatar/services/center/protos"
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/copier"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func GetById(
	ctx context.Context,
	db neo4j.Driver,
	routineRepository repository.RoutineRepository,
	input *pb.GetByIdRequest,
) (*pb.CreateRoutineResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	routineRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		Routine, err := routineRepository.GetById(ctx, input.Id)
		if err != nil {
			log.Error("error when write transaction, error: ", err)
			return nil, err
		}
		return Routine, nil
	})
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	routine := routineRaw.(*entity.CenterRoutine)
	var result pb.CreateRoutineResponse
	err = copier.Copy(&result, routine)
	if err != nil {
		return nil, err
	}
	result.StartDate = routine.StartDate.Format(time.RFC3339)
	result.EndDate = routine.EndDate.Format(time.RFC3339)
	result.CreatedAt = routine.CreatedAt.Format(time.RFC3339)
	if routine.UpdatedAt.IsZero() {
		result.UpdatedAt = ""
	} else {
		result.UpdatedAt = routine.UpdatedAt.Format(time.RFC3339)
	}
	return &result, nil
}
