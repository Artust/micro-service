package user_activity

import (
	"avatar/services/account_management/config"
	"avatar/services/account_management/domain/entity"
	"avatar/services/account_management/domain/repository"
	pb "avatar/services/account_management/protos"
	"context"

	"github.com/jinzhu/copier"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func Save(
	ctx context.Context,
	db neo4j.Driver,
	userActivityRepository repository.UserActivityRepository,
	input *pb.UserActivity,
) (*pb.UserActivity, error) {
	data := entity.UserActivity{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}

	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	activity, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		activity, err := userActivityRepository.Save(ctx, &data)
		if err != nil {
			return nil, err
		}
		return activity, nil
	})
	if err != nil {
		return nil, err
	}

	actUser := activity.(*entity.UserActivity)
	var result pb.UserActivity
	err = copier.Copy(&result, actUser)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
