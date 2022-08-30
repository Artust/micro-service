package ip_camera

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

func Create(
	ctx context.Context,
	db neo4j.Driver,
	ipCameraRepository repository.IpCameraRepository,
	input *pb.CreateIpCameraRequest,
) (*pb.CreateIpCameraResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := entity.IpCamera{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	createdIpCamera, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		ipCamera, err := ipCameraRepository.Create(ctx, &data)
		if err != nil {
			log.Error("error when create ip camera, error: ", err)
			return nil, err
		}
		return *ipCamera, nil
	})
	if err != nil {
		log.Error("error when create ip camera, error: ", err)
		return nil, err
	}
	ipCamera := createdIpCamera.(entity.IpCamera)
	var result pb.CreateIpCameraResponse
	err = copier.Copy(&result, ipCamera)
	if err != nil {
		return nil, err
	}
	result.CreatedAt = ipCamera.CreatedAt.Format(time.RFC3339)
	return &result, nil
}
