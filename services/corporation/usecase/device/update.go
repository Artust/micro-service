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

func Update(
	ctx context.Context,
	db neo4j.Driver,
	deviceRepository repository.DeviceRepository,
	input *pb.UpdateDeviceRequest,
) (*pb.CreateDeviceResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := entity.Device{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	deviceRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		device, err := deviceRepository.Update(ctx, data.Id, &data)
		if err != nil {
			log.Error("error when write transaction, error: ", err)
			return nil, err
		}
		return *device, nil
	})
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	device := deviceRaw.(entity.Device)
	var result pb.CreateDeviceResponse
	err = copier.Copy(&result, device)
	if err != nil {
		return nil, err
	}
	result.CreatedAt = device.CreatedAt.Format(time.RFC3339)
	result.UpdatedAt = device.UpdatedAt.Format(time.RFC3339)
	return &result, nil
}
