package account_role

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
	accountRoleRepository repository.AccountRoleRepository,
	input *pb.Id,
) (*pb.AccountRole, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	accountRole, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		role, err := accountRoleRepository.GetById(ctx, input.Id)
		if err != nil {
			return nil, err
		}
		return role, nil
	})
	if err != nil {
		return nil, err
	}

	role := accountRole.(*entity.AccountRole)
	var result pb.AccountRole
	err = copier.Copy(&result, role)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
