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

func Update(
	ctx context.Context,
	db neo4j.Driver,
	corporationRepository repository.CorporationRepository,
	input *pb.UpdateCorporationRequest,
) (*pb.CreateCorporationResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := entity.Corporation{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	corporationRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		corporation, err := corporationRepository.Update(ctx, data.Id, &data)
		if err != nil {
			log.Error("error when write transaction, error: ", err)
			return nil, err
		}
		return *corporation, nil
	})
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	corporation := corporationRaw.(entity.Corporation)
	var result pb.CreateCorporationResponse
	err = copier.Copy(&result, corporation)
	if err != nil {
		return nil, err
	}
	result.CreatedAt = corporation.CreatedAt.Format(time.RFC3339)
	result.UpdatedAt = corporation.UpdatedAt.Format(time.RFC3339)
	return &result, nil
}
