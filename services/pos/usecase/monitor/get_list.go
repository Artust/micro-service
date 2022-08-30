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

func GetList(
	ctx context.Context,
	db neo4j.Driver,
	monitorRepository repository.MonitorRepository,
	input *pb.GetListMonitorRequest,
) (*pb.GetListMonitorResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := entity.GetListMonitorOption{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	monitorsRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		monitors, err := monitorRepository.GetList(ctx, data)
		if err != nil {
			log.Error("error when get list monitor, error: ", err)
			return nil, err
		}
		return monitors, nil
	})
	if err != nil {
		log.Error("error when get list monitor, error: ", err)
		return nil, err
	}
	monitors := monitorsRaw.([]*entity.Monitor)
	var results pb.GetListMonitorResponse
	results.GetListMonitorResponse = make([]*pb.CreateMonitorResponse, 0)
	for _, monitor := range monitors {
		var response pb.CreateMonitorResponse
		err = copier.Copy(&response, monitor)
		if err != nil {
			return nil, err
		}
		response.CreatedAt = monitor.CreatedAt.Format(time.RFC3339)
		if monitor.UpdatedAt.IsZero() {
			response.UpdatedAt = ""
		} else {
			response.UpdatedAt = monitor.UpdatedAt.Format(time.RFC3339)
		}
		results.GetListMonitorResponse = append(results.GetListMonitorResponse, &response)
	}
	return &results, nil
}
