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

func GetList(
	ctx context.Context,
	db neo4j.Driver,
	routineRepository repository.RoutineRepository,
	input *pb.GetListRoutineRequest,
) (*pb.GetListRoutineResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := entity.GetListRoutineOption{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	routinesRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		routines, err := routineRepository.GetList(ctx, data)
		if err != nil {
			log.Error("error when get list routine, error: ", err)
			return nil, err
		}
		return routines, nil
	})
	if err != nil {
		log.Error("error when get list routine, error: ", err)
		return nil, err
	}
	routines := routinesRaw.([]*entity.PosRoutine)
	var results pb.GetListRoutineResponse
	results.GetListRoutineResponse = make([]*pb.CreateRoutineResponse, 0)
	for _, routine := range routines {
		var response pb.CreateRoutineResponse
		err = copier.Copy(&response, routine)
		response.StartDate = routine.StartDate.Format(time.RFC3339)
		response.EndDate = routine.EndDate.Format(time.RFC3339)
		if err != nil {
			return nil, err
		}
		response.CreatedAt = routine.CreatedAt.Format(time.RFC3339)
		if routine.UpdatedAt.IsZero() {
			response.UpdatedAt = ""
		} else {
			response.UpdatedAt = routine.UpdatedAt.Format(time.RFC3339)
		}
		results.GetListRoutineResponse = append(results.GetListRoutineResponse, &response)
	}

	return &results, nil
}
