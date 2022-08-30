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

func GetById(
	ctx context.Context,
	db neo4j.Driver,
	routineCategoryRepository repository.RoutineCategoryRepository,
	input *pb.GetByIdRequest,
) (*pb.CreateRoutineCategoryResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	routineCategoryRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		routineCategory, err := routineCategoryRepository.GetById(ctx, input.Id)
		if err != nil {
			log.Error("error when get routine category, error: ", err)
			return nil, err
		}
		return routineCategory, nil
	})
	if err != nil {
		log.Error("error when get routine category, error: ", err)
		return nil, err
	}
	routineCategory := routineCategoryRaw.(*entity.RoutineCategory)
	var result pb.CreateRoutineCategoryResponse
	err = copier.Copy(&result, routineCategory)
	if err != nil {
		return nil, err
	}
	result.CreatedAt = routineCategory.CreatedAt.Format(time.RFC3339)
	if routineCategory.UpdatedAt.IsZero() {
		result.UpdatedAt = ""
	} else {
		result.UpdatedAt = routineCategory.UpdatedAt.Format(time.RFC3339)
	}
	return &result, nil
}
