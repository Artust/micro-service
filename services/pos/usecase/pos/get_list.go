package pos

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
	posRepository repository.PosRepository,
	input *pb.GetListPosRequest,
) (*pb.GetListPosResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := entity.GetListPosOption{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	possRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		poss, err := posRepository.GetList(ctx, data)
		if err != nil {
			log.Error("error when get list pos, error: ", err)
			return nil, err
		}
		return poss, nil
	})
	if err != nil {
		log.Error("error when get list pos, error: ", err)
		return nil, err
	}
	poss := possRaw.([]*entity.Pos)
	var results pb.GetListPosResponse
	results.GetListPosResponse = make([]*pb.CreatePosResponse, 0)
	for _, pos := range poss {
		var response pb.CreatePosResponse
		err = copier.Copy(&response, pos)
		if err != nil {
			return nil, err
		}
		response.CreatedAt = pos.CreatedAt.Format(time.RFC3339)
		if pos.UpdatedAt.IsZero() {
			response.UpdatedAt = ""
		} else {
			response.UpdatedAt = pos.UpdatedAt.Format(time.RFC3339)
		}
		results.GetListPosResponse = append(results.GetListPosResponse, &response)
	}
	return &results, nil
}
