package routine

import (
	"avatar/services/center/config"
	"avatar/services/center/domain/entity"
	"avatar/services/center/domain/repository"
	pb "avatar/services/center/protos"
	"context"
	"time"

	"github.com/jinzhu/copier"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func GetListByCategory(
	ctx context.Context,
	db neo4j.Driver,
	routineRepository repository.RoutineRepository,
	routineCategoryRepository repository.RoutineCategoryRepository,
	input *pb.GetListRoutineByCategoryRequest,
) (*pb.GetListRoutineByCategoryResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := entity.GetListRoutineOption{}
	dataCategory := entity.GetListRoutineCategoryOption{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	var categories []*entity.RoutineCategory
	routinesRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		categories, err = routineCategoryRepository.GetList(ctx, dataCategory)
		if err != nil {
			log.Error("error when get list routine category, error: ", err)
			return nil, err
		}
		for _, val := range categories {
			data.CategoryIds = append(data.CategoryIds, val.Id)
		}
		routines, err := routineRepository.GetListRoutineCategory(ctx, data)
		if err != nil {
			log.Error("error when get list routine by category, error: ", err)
			return nil, err
		}
		return routines, nil
	})
	if err != nil {
		log.Error("error when get list routine by category, error: ", err)
		return nil, err
	}
	routines := routinesRaw.([]*entity.CenterRoutine)
	var output pb.GetListRoutineByCategoryResponse
	for _, valCategory := range categories {
		var result pb.ListRoutineByCategory
		for _, valRoutine := range routines {
			if valRoutine.CategoryId == valCategory.Id {
				var updatedAtCategory, updatedAtRoutine string
				if valCategory.UpdatedAt.IsZero() {
					updatedAtCategory = ""
				} else {
					updatedAtCategory = valCategory.UpdatedAt.Format(time.RFC3339)
				}
				result.Category = &pb.CreateRoutineCategoryResponse{
					Id:        valCategory.Id,
					Name:      valCategory.Name,
					CreatedAt: valCategory.CreatedAt.Format(time.RFC3339),
					UpdatedAt: updatedAtCategory,
				}
				if valRoutine.UpdatedAt.IsZero() {
					updatedAtRoutine = ""
				} else {
					updatedAtRoutine = valRoutine.UpdatedAt.Format(time.RFC3339)
				}
				result.Routine = append(result.Routine, &pb.CreateRoutineResponse{
					Id:            valRoutine.Id,
					Name:          valRoutine.Name,
					Detail:        valRoutine.Detail,
					AnimationFile: valRoutine.AnimationFile,
					SoundFile:     valRoutine.SoundFile,
					ImageFile:     valRoutine.ImageFile,
					CategoryId:    valRoutine.CategoryId,
					StartDate:     valRoutine.StartDate.Format(time.RFC3339),
					EndDate:       valRoutine.EndDate.Format(time.RFC3339),
					CreatedAt:     valRoutine.CreatedAt.Format(time.RFC3339),
					UpdatedAt:     updatedAtRoutine,
				})
			}
		}
		if result.Routine != nil {
			output.GetListRoutineByCategoryResponse = append(output.GetListRoutineByCategoryResponse, &result)
		}
	}
	return &output, nil
}
