package account

import (
	"avatar/services/account_management/config"
	"avatar/services/account_management/domain/entity"
	"avatar/services/account_management/domain/repository"
	pb "avatar/services/account_management/protos"
	"context"

	"github.com/jinzhu/copier"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func GetById(
	ctx context.Context,
	db neo4j.Driver,
	accountRepository repository.AccountRepository,
	input *pb.Id,
) (*pb.Account, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	newInfo, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		acc, err := accountRepository.GetById(ctx, input.Id)
		if err != nil {
			return nil, err
		}
		return acc, nil
	})
	if err != nil {
		return nil, err
	}
	acc := newInfo.(*entity.Account)
	var result pb.Account
	err = copier.Copy(&result, acc)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
