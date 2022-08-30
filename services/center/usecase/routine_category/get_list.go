package routine_category

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
	routineCategoryRepository repository.RoutineCategoryRepository,
	input *pb.GetListRoutineCategoryRequest,
) (*pb.GetListRoutineCategoryResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := entity.GetListRoutineCategoryOption{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	routineCategorysRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		routineCategorys, err := routineCategoryRepository.GetList(ctx, data)
		if err != nil {
			log.Error("error when write transaction, error: ", err)
			return nil, err
		}
		return routineCategorys, nil
	})
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	routineCategorys := routineCategorysRaw.([]*entity.RoutineCategory)
	var results pb.GetListRoutineCategoryResponse
	results.GetListRoutineCategoryResponse = make([]*pb.CreateRoutineResponse, 0)
	for _, routineCategorys := range routineCategorys {
		var response pb.CreateRoutineResponse
		err = copier.Copy(&response, routineCategorys)
		if err != nil {
			return nil, err
		}
		response.CreatedAt = routineCategorys.CreatedAt.Format(time.RFC3339)
		if routineCategorys.UpdatedAt.IsZero() {
			response.UpdatedAt = ""
		} else {
			response.UpdatedAt = routineCategorys.UpdatedAt.Format(time.RFC3339)
		}
		results.GetListRoutineCategoryResponse = append(results.GetListRoutineCategoryResponse, &response)
	}
	return &results, nil
}
