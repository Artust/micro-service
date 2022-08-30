package corporation

import (
	"avatar/services/corporation/config"
	"avatar/services/corporation/domain/entity"
	"avatar/services/corporation/domain/repository"
	pb "avatar/services/corporation/protos"
	"context"
	"time"

	"github.com/jinzhu/copier"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func GetList(
	ctx context.Context,
	db neo4j.Driver,
	corporationRepository repository.CorporationRepository,
	input *pb.GetListCorporationRequest,
) (*pb.GetListCorporationResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := entity.GetListCorporationOption{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	corporationsRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		corporations, err := corporationRepository.GetList(ctx, data)
		if err != nil {
			log.Error("error when write transaction, error: ", err)
			return nil, err
		}
		return corporations, nil
	})
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	corporations := corporationsRaw.([]*entity.Corporation)
	var results pb.GetListCorporationResponse
	results.GetListCorporationResponse = make([]*pb.CreateCorporationResponse, 0)
	for _, corporation := range corporations {
		var response pb.CreateCorporationResponse
		err = copier.Copy(&response, corporation)
		if err != nil {
			return nil, err
		}
		response.CreatedAt = corporation.CreatedAt.Format(time.RFC3339)
		if corporation.UpdatedAt.IsZero() {
			response.UpdatedAt = ""
		} else {
			response.UpdatedAt = corporation.UpdatedAt.Format(time.RFC3339)
		}
		results.GetListCorporationResponse = append(results.GetListCorporationResponse, &response)
	}
	return &results, nil
}
