package routine

import (
	"avatar/services/pos/config"
	"avatar/services/pos/domain/entity"
	"avatar/services/pos/domain/repository"
	pb "avatar/services/pos/protos"
	"context"

	"github.com/jinzhu/copier"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func CreateMany(
	ctx context.Context,
	db neo4j.Driver,
	routineRepository repository.RoutineRepository,
	input *pb.CreateManyRoutineRequest,
) (*pb.CreateManyRoutineResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := []entity.PosRoutine{}
	err := copier.Copy(&data, input.Routines)
	if err != nil {
		return nil, err
	}
	createdRoutine, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		routine, err := routineRepository.CreateMany(ctx, &data)
		if err != nil {
			log.Error("error when create routine, error: ", err)
			return nil, err
		}
		return *routine, nil
	})
	if err != nil {
		log.Error("error when create routine, error: ", err)
		return nil, err
	}
	routine := createdRoutine.([]entity.PosRoutine)
	var result pb.CreateManyRoutineResponse
	err = copier.Copy(&result.Routines, routine)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
