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

func GetList(
	ctx context.Context,
	db neo4j.Driver,
	centerRepository repository.CenterRepository,
	input *pb.GetListCenterRequest,
) (*pb.GetListCenterResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := entity.GetListCenterOption{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	CentersRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		Centers, err := centerRepository.GetList(ctx, data)
		if err != nil {
			log.Error("error when write transaction, error: ", err)
			return nil, err
		}
		return Centers, nil
	})
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	centers := CentersRaw.([]*entity.Center)
	var results pb.GetListCenterResponse
	results.ListCenter = make([]*pb.CreateCenterResponse, 0)
	for _, center := range centers {
		var response pb.CreateCenterResponse
		err = copier.Copy(&response, center)
		if err != nil {
			return nil, err
		}
		response.CreatedAt = center.CreatedAt.Format(time.RFC3339)
		if center.UpdatedAt.IsZero() {
			response.UpdatedAt = ""
		} else {
			response.UpdatedAt = center.UpdatedAt.Format(time.RFC3339)
		}
		results.ListCenter = append(results.ListCenter, &response)
	}
	return &results, nil
}
