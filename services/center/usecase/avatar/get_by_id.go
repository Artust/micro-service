package avatar

import (
	"avatar/services/center/config"
	"avatar/services/center/domain/repository"
	"avatar/services/center/domain/entity"
	pb "avatar/services/center/protos"
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/copier"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func GetById(
	ctx context.Context,
	db neo4j.Driver,
	avatarRepository repository.AvatarRepository,
	input *pb.GetByIdRequest,
) (*pb.CreateAvatarResponse, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	avatarRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		avatar, err := avatarRepository.GetById(ctx, input.Id)
		if err != nil {
			log.Error("error when write transaction, error: ", err)
			return nil, err
		}
		return avatar, nil
	})
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return nil, err
	}
	avatar := avatarRaw.(*entity.Avatar)
	var result pb.CreateAvatarResponse
	err = copier.Copy(&result, avatar)
	if err != nil {
		return nil, err
	}
	result.StartDate = avatar.StartDate.Format(time.RFC3339)
	result.EndDate = avatar.EndDate.Format(time.RFC3339)
	result.CreatedAt = avatar.CreatedAt.Format(time.RFC3339)
	if avatar.UpdatedAt.IsZero() {
		result.UpdatedAt = ""
	} else {
		result.UpdatedAt = avatar.UpdatedAt.Format(time.RFC3339)
	}
	return &result, nil
}
