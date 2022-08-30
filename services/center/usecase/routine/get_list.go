package routine

import (
	"avatar/services/center/config"
	"avatar/services/center/domain/repository"
	"avatar/services/center/domain/entity"
	pb "avatar/services/center/protos"
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/copier"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
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
			log.Error("error when write transaction, error: ", err)
			return nil, err
		}
		return routines, nil
	})
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	Routines := routinesRaw.([]*entity.CenterRoutine)
	var results pb.GetListRoutineResponse
	results.ListRoutine = make([]*pb.CreateRoutineResponse, 0)
	for _, Routines := range Routines {
		var response pb.CreateRoutineResponse
		err = copier.Copy(&response, Routines)
		if err != nil {
			return nil, err
		}
		response.StartDate = Routines.StartDate.Format(time.RFC3339)
		response.EndDate = Routines.EndDate.Format(time.RFC3339)
		response.CreatedAt = Routines.CreatedAt.Format(time.RFC3339)
		if Routines.UpdatedAt.IsZero() {
			response.UpdatedAt = ""
		} else {
			response.UpdatedAt = Routines.UpdatedAt.Format(time.RFC3339)
		}
		results.ListRoutine = append(results.ListRoutine, &response)
	}
	return &results, nil
}
