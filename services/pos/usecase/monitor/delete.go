package monitor

import (
	"avatar/services/pos/config"
	"avatar/services/pos/domain/repository"
	pb "avatar/services/pos/protos"
	"context"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func Delete(
	ctx context.Context,
	db neo4j.Driver,
	monitorRepository repository.MonitorRepository,
	input *pb.DeleteByIdRequest,
) (*pb.DeleteResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	rowsAffectedRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		rowsAffected, err := monitorRepository.Delete(ctx, input.Id)
		if err != nil {
			log.Error("error when delete monitor, error: ", err)
			return nil, err
		}
		return rowsAffected, nil
	})
	if err != nil {
		log.Error("error when delete monitor, error: ", err)
		return nil, err
	}
	return &pb.DeleteResponse{
		RowsAffected: rowsAffectedRaw.(int64),
	}, nil
}
