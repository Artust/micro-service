package monitor

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
	monitorRepository repository.MonitorRepository,
	input *pb.GetByIdRequest,
) (*pb.CreateMonitorResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	monitorRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		monitor, err := monitorRepository.GetById(ctx, input.Id)
		if err != nil {
			log.Error("error when get monitor, error: ", err)
			return nil, err
		}
		return monitor, nil
	})
	if err != nil {
		log.Error("error when get monitor, error: ", err)
		return nil, err
	}
	monitor := monitorRaw.(*entity.Monitor)
	var result pb.CreateMonitorResponse
	err = copier.Copy(&result, monitor)
	if err != nil {
		return nil, err
	}
	result.CreatedAt = monitor.CreatedAt.Format(time.RFC3339)
	if monitor.UpdatedAt.IsZero() {
		result.UpdatedAt = ""
	} else {
		result.UpdatedAt = monitor.UpdatedAt.Format(time.RFC3339)
	}
	return &result, nil
}
