package routine

import (
	"avatar/services/pos/config"
	"avatar/services/pos/domain/entity"
	"avatar/services/pos/domain/repository"
	pb "avatar/services/pos/protos"
	"context"
	"time"

	"github.com/jinzhu/copier"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func Update(
	ctx context.Context,
	db neo4j.Driver,
	routineRepository repository.RoutineRepository,
	input *pb.UpdateRoutineRequest,
) (*pb.CreateRoutineResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := entity.PosRoutine{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	data.StartDate, err = time.Parse(time.RFC3339, input.StartDate)
	if err != nil {
		return nil, err
	}
	data.EndDate, err = time.Parse(time.RFC3339, input.EndDate)
	if err != nil {
		return nil, err
	}
	routineRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		routine, err := routineRepository.Update(ctx, data.Id, &data)
		if err != nil {
			log.Error("error when update routine, error: ", err)
			return nil, err
		}
		return *routine, nil
	})

	if err != nil {
		log.Error("error when update routine, error: ", err)
		return nil, err
	}
	routine := routineRaw.(entity.PosRoutine)
	var result pb.CreateRoutineResponse
	err = copier.Copy(&result, routine)
	if err != nil {
		return nil, err
	}
	result.StartDate = routine.StartDate.Format(time.RFC3339)
	result.EndDate = routine.EndDate.Format(time.RFC3339)
	result.CreatedAt = routine.CreatedAt.Format(time.RFC3339)
	result.UpdatedAt = routine.UpdatedAt.Format(time.RFC3339)
	return &result, nil
}
