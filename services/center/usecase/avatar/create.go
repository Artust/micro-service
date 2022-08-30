package avatar

import (
	"avatar/services/center/config"
	"avatar/services/center/domain/entity"
	"avatar/services/center/domain/repository"
	pb "avatar/services/center/protos"
	"context"
	"time"

	"github.com/jinzhu/copier"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

func Create(
	ctx context.Context,
	db neo4j.Driver,
	avatarRepository repository.AvatarRepository,
	input *pb.CreateAvatarRequest,
) (*pb.CreateAvatarResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	data := entity.Avatar{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	data.StartDate, err = time.Parse(time.RFC3339, input.StartDate)
	if err != nil {
		return nil, err
	}
	data.EndDate, err = time.Parse(time.RFC3339, input.EndDate)
	if err != nil {
		return nil, err
	}
	avatarRow, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		avatar, err := avatarRepository.Create(ctx, &data)
		if err != nil {
			log.Error("error when write transaction, error: ", err)
			return nil, err
		}
		return *avatar, nil
	})
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	avatar := avatarRow.(entity.Avatar)
	var result pb.CreateAvatarResponse
	err = copier.Copy(&result, avatar)
	if err != nil {
		return nil, err
	}
	result.StartDate = avatar.StartDate.Format(time.RFC3339)
	result.EndDate = avatar.EndDate.Format(time.RFC3339)
	result.CreatedAt = avatar.CreatedAt.Format(time.RFC3339)
	return &result, nil
}
