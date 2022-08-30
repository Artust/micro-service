package user_activity

import (
	"avatar/services/account_management/config"
	"avatar/services/account_management/domain/entity"
	"avatar/services/account_management/domain/repository"
	pb "avatar/services/account_management/protos"
	"context"
	"time"

	"github.com/jinzhu/copier"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func Get(
	ctx context.Context,
	db neo4j.Driver,
	userActivityRepository repository.UserActivityRepository,
	input *pb.AccountId,
) (*pb.UserActivityList, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	activities, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		acc, errReset := userActivityRepository.Get(ctx, input.AccountId)
		if errReset != nil {
			return nil, errReset
		}
		return acc, nil
	})
	if err != nil {
		return nil, err
	}
	userActivityList := activities.([]*entity.UserActivity)
	var results pb.UserActivityList
	results.Activities = make([]*pb.UserActivity, 0)
	for _, acc := range userActivityList {
		var response pb.UserActivity
		err = copier.Copy(&response, acc)
		if err != nil {
			return nil, err
		}
		response.CreatedAt = acc.CreatedAt.Format(time.RFC3339)
		if acc.UpdatedAt.IsZero() {
			response.UpdatedAt = ""
		} else {
			response.UpdatedAt = acc.UpdatedAt.Format(time.RFC3339)
		}
		results.Activities = append(results.Activities, &response)
	}
	return &results, nil
}
