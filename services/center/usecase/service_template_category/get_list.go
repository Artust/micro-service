package service_template_category

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
	serviceTemplateCategoryRepository repository.ServiceTemplateCategoryRepository,
	input *pb.GetListServiceTemplateCategoryRequest,
) (*pb.GetListServiceTemplateCategoryResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := entity.GetListServiceTemplateCategoryOption{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	ServiceTemplateCategorysRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		ServiceTemplateCategorys, err := serviceTemplateCategoryRepository.GetList(ctx, data)
		if err != nil {
			log.Error("error when write transaction, error: ", err)
			return nil, err
		}
		return ServiceTemplateCategorys, nil
	})
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	ServiceTemplateCategorys := ServiceTemplateCategorysRaw.([]*entity.ServiceTemplateCategory)
	var results pb.GetListServiceTemplateCategoryResponse
	results.GetListServiceTemplateCategoryResponse = make([]*pb.CreateServiceTemplateCategoryResponse, 0)
	for _, ServiceTemplateCategorys := range ServiceTemplateCategorys {
		var response pb.CreateServiceTemplateCategoryResponse
		err = copier.Copy(&response, ServiceTemplateCategorys)
		if err != nil {
			return nil, err
		}
		response.CreatedAt = ServiceTemplateCategorys.CreatedAt.Format(time.RFC3339)
		if ServiceTemplateCategorys.UpdatedAt.IsZero() {
			response.UpdatedAt = ""
		} else {
			response.UpdatedAt = ServiceTemplateCategorys.UpdatedAt.Format(time.RFC3339)
		}
		results.GetListServiceTemplateCategoryResponse = append(results.GetListServiceTemplateCategoryResponse, &response)
	}
	return &results, nil
}
