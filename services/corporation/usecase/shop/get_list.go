package shop

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
	shopRepository repository.ShopRepository,
	input *pb.GetListShopRequest,
) (*pb.GetListShopResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := entity.GetListShopOption{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	shopsRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		shops, err := shopRepository.GetList(ctx, data)
		if err != nil {
			log.Error("error when write transaction, error: ", err)
			return nil, err
		}
		return shops, nil
	})
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	shops := shopsRaw.([]*entity.Shop)
	var results pb.GetListShopResponse
	results.GetListShopResponse = make([]*pb.CreateShopResponse, 0)
	for _, shop := range shops {
		var response pb.CreateShopResponse
		err = copier.Copy(&response, shop)
		if err != nil {
			return nil, err
		}
		response.CreatedAt = shop.CreatedAt.Format(time.RFC3339)
		if shop.UpdatedAt.IsZero() {
			response.UpdatedAt = ""
		} else {
			response.UpdatedAt = shop.UpdatedAt.Format(time.RFC3339)
		}
		results.GetListShopResponse = append(results.GetListShopResponse, &response)
	}
	return &results, nil
}
