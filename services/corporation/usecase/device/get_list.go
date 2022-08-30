package device

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
	deviceRepository repository.DeviceRepository,
	input *pb.GetListDeviceRequest,
) (*pb.GetListDeviceResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := entity.GetListDeviceOption{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	devicesRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		devices, err := deviceRepository.GetList(ctx, data)
		if err != nil {
			log.Error("error when write transaction, error: ", err)
			return nil, err
		}
		return devices, nil
	})
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	devices := devicesRaw.([]*entity.Device)
	var results pb.GetListDeviceResponse
	results.GetListDeviceResponse = make([]*pb.CreateDeviceResponse, 0)
	for _, device := range devices {
		var response pb.CreateDeviceResponse
		err = copier.Copy(&response, device)
		if err != nil {
			return nil, err
		}
		response.CreatedAt = device.CreatedAt.Format(time.RFC3339)
		if device.UpdatedAt.IsZero() {
			response.UpdatedAt = ""
		} else {
			response.UpdatedAt = device.UpdatedAt.Format(time.RFC3339)
		}
		results.GetListDeviceResponse = append(results.GetListDeviceResponse, &response)
	}
	return &results, nil
}
