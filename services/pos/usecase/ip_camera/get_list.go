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

func GetList(
	ctx context.Context,
	db neo4j.Driver,
	ipCameraRepository repository.IpCameraRepository,
	input *pb.GetListIpCameraRequest,
) (*pb.GetListIpCameraResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := entity.GetListIpCameraOption{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	ipCamerasRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		ipCameras, err := ipCameraRepository.GetList(ctx, data)
		if err != nil {
			log.Error("error when get list ip camera, error: ", err)
			return nil, err
		}
		return ipCameras, nil
	})
	if err != nil {
		log.Error("error when get list ip camera, error: ", err)
		return nil, err
	}
	ipCameras := ipCamerasRaw.([]*entity.IpCamera)
	var results pb.GetListIpCameraResponse
	results.GetListIpCameraResponse = make([]*pb.CreateIpCameraResponse, 0)
	for _, ipCamera := range ipCameras {
		var response pb.CreateIpCameraResponse
		err = copier.Copy(&response, ipCamera)
		if err != nil {
			return nil, err
		}
		response.CreatedAt = ipCamera.CreatedAt.Format(time.RFC3339)
		if ipCamera.UpdatedAt.IsZero() {
			response.UpdatedAt = ""
		} else {
			response.UpdatedAt = ipCamera.UpdatedAt.Format(time.RFC3339)
		}
		results.GetListIpCameraResponse = append(results.GetListIpCameraResponse, &response)
	}
	return &results, nil
}
