package service_template

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
	serviceTemplateRepository repository.ServiceTemplateRepository,
	input *pb.GetListServiceTemplateRequest,
) (*pb.GetListServiceTemplateResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := entity.GetListServiceTemplateOption{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	serviceTemplatesRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		serviceTemplates, err := serviceTemplateRepository.GetList(ctx, data)
		if err != nil {
			log.Error("error when write transaction, error: ", err)
			return nil, err
		}
		return serviceTemplates, nil
	})
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	serviceTemplates := serviceTemplatesRaw.([]*entity.ServiceTemplate)
	var results pb.GetListServiceTemplateResponse
	results.ListServiceTemplate = make([]*pb.CreateServiceTemplateResponse, 0)
	for _, serviceTemplates := range serviceTemplates {
		var response pb.CreateServiceTemplateResponse
		err = copier.Copy(&response, serviceTemplates)
		if err != nil {
			return nil, err
		}
		response.CreatedAt = serviceTemplates.CreatedAt.Format(time.RFC3339)
		if serviceTemplates.UpdatedAt.IsZero() {
			response.UpdatedAt = ""
		} else {
			response.UpdatedAt = serviceTemplates.UpdatedAt.Format(time.RFC3339)
		}
		results.ListServiceTemplate = append(results.ListServiceTemplate, &response)
	}
	return &results, nil
}
