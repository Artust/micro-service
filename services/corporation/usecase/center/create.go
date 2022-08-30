package center

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

func Create(
	ctx context.Context,
	db neo4j.Driver,
	centerRepository repository.CenterRepository,
	input *pb.CreateCenterRequest,
) (*pb.CreateCenterResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := entity.Center{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	createdCenter, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		center, err := centerRepository.Create(ctx, &data)
		if err != nil {
			log.Error("error when write transaction, error: ", err)
			return nil, err
		}
		return *center, nil
	})
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	center := createdCenter.(entity.Center)
	var result pb.CreateCenterResponse
	err = copier.Copy(&result, center)
	if err != nil {
		return nil, err
	}
	result.CreatedAt = center.CreatedAt.Format(time.RFC3339)
	return &result, nil
}
