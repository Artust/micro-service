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

func Update(
	ctx context.Context,
	db neo4j.Driver,
	serviceTemplateRepository repository.ServiceTemplateRepository,
	input *pb.UpdateServiceTemplateRequest,
) (*pb.CreateServiceTemplateResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := entity.ServiceTemplate{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	serviceTemplateRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		serviceTemplate, err := serviceTemplateRepository.Update(ctx, data.Id, &data)
		if err != nil {
			log.Error("error when write transaction, error: ", err)
			return nil, err
		}
		return *serviceTemplate, nil
	})
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	serviceTemplate := serviceTemplateRaw.(entity.ServiceTemplate)
	var result pb.CreateServiceTemplateResponse
	err = copier.Copy(&result, serviceTemplate)
	if err != nil {
		return nil, err
	}
	result.CreatedAt = serviceTemplate.CreatedAt.Format(time.RFC3339)
	result.UpdatedAt = serviceTemplate.UpdatedAt.Format(time.RFC3339)
	return &result, nil
}
