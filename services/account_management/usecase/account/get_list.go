package account

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

func GetList(
	ctx context.Context,
	db neo4j.Driver,
	accountRepository repository.AccountRepository,
	input *pb.GetListAccountRequest,
) (*pb.AccountList, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	accounts, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		acc, errReset := accountRepository.GetList(ctx, input)
		if errReset != nil {
			return nil, errReset
		}
		return acc, nil
	})
	if err != nil {
		return nil, err
	}
	accountList := accounts.([]*entity.Account)
	var results pb.AccountList
	results.Accounts = make([]*pb.Account, 0)
	for _, acc := range accountList {
		var response pb.Account
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
		results.Accounts = append(results.Accounts, &response)
	}
	return &results, nil
}
