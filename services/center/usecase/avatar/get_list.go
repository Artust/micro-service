package avatar

import (
	"avatar/services/center/config"
	"avatar/services/center/domain/entity"
	"avatar/services/center/domain/repository"
	pb "avatar/services/center/protos"
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/copier"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func GetList(
	ctx context.Context,
	db neo4j.Driver,
	avatarRepository repository.AvatarRepository,
	input *pb.GetListAvatarRequest,
) (*pb.GetListAvatarResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := entity.GetListAvatarOption{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	avatarsRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		avatars, err := avatarRepository.GetList(ctx, data)
		if err != nil {
			log.Error("error when write transaction, error: ", err)
			return nil, err
		}
		return avatars, nil
	})
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	avatars := avatarsRaw.([]*entity.Avatar)
	var results pb.GetListAvatarResponse
	results.Avatars = make([]*pb.CreateAvatarResponse, 0)
	for _, avatars := range avatars {
		var response pb.CreateAvatarResponse
		err = copier.Copy(&response, avatars)
		if err != nil {
			return nil, err
		}
		response.StartDate = avatars.StartDate.Format(time.RFC3339)
		response.EndDate = avatars.EndDate.Format(time.RFC3339)
		response.CreatedAt = avatars.CreatedAt.Format(time.RFC3339)
		if avatars.UpdatedAt.IsZero() {
			response.UpdatedAt = ""
		} else {
			response.UpdatedAt = avatars.UpdatedAt.Format(time.RFC3339)
		}
		results.Avatars = append(results.Avatars, &response)
	}
	return &results, nil
}
