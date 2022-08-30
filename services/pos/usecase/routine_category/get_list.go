package routine_category

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
			log.Error("error when get list routine category, error: ", err)
			return nil, err
		}
		return routineCategorys, nil
	})
	if err != nil {
		log.Error("error when get list routine category, error: ", err)
		return nil, err
	}
	routineCategorys := routineCategorysRaw.([]*entity.RoutineCategory)
	var results pb.GetListRoutineCategoryResponse
	results.GetListRoutineCategoryResponse = make([]*pb.CreateRoutineCategoryResponse, 0)
	for _, routineCategory := range routineCategorys {
		var response pb.CreateRoutineCategoryResponse
		err = copier.Copy(&response, routineCategory)
		if err != nil {
			return nil, err
		}
		response.CreatedAt = routineCategory.CreatedAt.Format(time.RFC3339)
		if routineCategory.UpdatedAt.IsZero() {
			response.UpdatedAt = ""
		} else {
			response.UpdatedAt = routineCategory.UpdatedAt.Format(time.RFC3339)
		}
		results.GetListRoutineCategoryResponse = append(results.GetListRoutineCategoryResponse, &response)
	}
	return &results, nil
}
