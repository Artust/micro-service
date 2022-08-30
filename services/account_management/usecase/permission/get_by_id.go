package permission

import (
	"avatar/services/account_management/config"
	"avatar/services/account_management/domain/entity"
	"avatar/services/account_management/domain/repository"
	pb "avatar/services/account_management/protos"
	"context"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func GetById(
	ctx context.Context,
	db neo4j.Driver,
	permissionRepository repository.PermissionRepository,
	input *pb.Id,
) (*pb.Permission, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	permission, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		per, err := permissionRepository.GetById(ctx, input.Id)
		if err != nil {
			return nil, err
		}
		return per, nil
	})
	if err != nil {
		return nil, err
	}

	per := permission.(*entity.Permission)
	fmt.Println("per: ", per)
	var result pb.Permission
	err = copier.Copy(&result, per)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
