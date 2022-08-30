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

func GetById(
	ctx context.Context,
	db neo4j.Driver,
	ipCameraRepository repository.IpCameraRepository,
	input *pb.GetByIdRequest,
) (*pb.CreateIpCameraResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	ipCameraRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		ipCamera, err := ipCameraRepository.GetById(ctx, input.Id)
		if err != nil {
			log.Error("error when get ip camera, error: ", err)
			return nil, err
		}
		return ipCamera, nil
	})
	if err != nil {
		log.Error("error when get ip camera, error: ", err)
		return nil, err
	}
	ipCamera := ipCameraRaw.(*entity.IpCamera)
	var result pb.CreateIpCameraResponse
	err = copier.Copy(&result, ipCamera)
	if err != nil {
		return nil, err
	}
	result.CreatedAt = ipCamera.CreatedAt.Format(time.RFC3339)
	if ipCamera.UpdatedAt.IsZero() {
		result.UpdatedAt = ""
	} else {
		result.UpdatedAt = ipCamera.UpdatedAt.Format(time.RFC3339)
	}
	return &result, nil
}
