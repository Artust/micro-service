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

func Create(
	ctx context.Context,
	db neo4j.Driver,
	monitorRepository repository.MonitorRepository,
	input *pb.CreateMonitorRequest,
) (*pb.CreateMonitorResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := entity.Monitor{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	createdMonitor, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		monitor, err := monitorRepository.Create(ctx, &data)
		if err != nil {
			log.Error("error when create monitor, error: ", err)
			return nil, err
		}
		return *monitor, nil
	})
	if err != nil {
		log.Error("error when create monitor, error: ", err)
		return nil, err
	}
	monitor := createdMonitor.(entity.Monitor)
	var result pb.CreateMonitorResponse
	err = copier.Copy(&result, monitor)
	if err != nil {
		return nil, err
	}
	result.CreatedAt = monitor.CreatedAt.Format(time.RFC3339)
	return &result, nil
}
