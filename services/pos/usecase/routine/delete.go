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

func Delete(
	ctx context.Context,
	db neo4j.Driver,
	routineRepository repository.RoutineRepository,
	input *pb.DeleteRoutineRequest,
) (*pb.DeleteResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	deleteRoutineOption := entity.DeleteRoutineOption{}
	err := copier.Copy(&deleteRoutineOption, input)
	if err != nil {
		return nil, err
	}
	rowsAffectedRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		rowsAffected, err := routineRepository.Delete(ctx, &deleteRoutineOption)
		if err != nil {
			log.Error("error when delete routine, error: ", err)
			return nil, err
		}
		return rowsAffected, nil
	})
	if err != nil {
		log.Error("error when delete routine, error: ", err)
		return nil, err
	}
	return &pb.DeleteResponse{
		RowsAffected: rowsAffectedRaw.(int64),
	}, nil
}
