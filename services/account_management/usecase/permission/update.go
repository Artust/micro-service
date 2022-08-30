package permission

import (
	"avatar/services/account_management/config"
	"avatar/services/account_management/domain/entity"
	"avatar/services/account_management/domain/repository"
	pb "avatar/services/account_management/protos"
	"context"

	"github.com/jinzhu/copier"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func Update(
	ctx context.Context,
	db neo4j.Driver,
	permissionRepository repository.PermissionRepository,
	input *pb.UpdatePermissionRequest,
) (*pb.Permission, error) {
	data := entity.Permission{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}

	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	newInfo, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		per, err := permissionRepository.Update(ctx, &data)
		if err != nil {
			return nil, err
		}
		return per, nil
	})
	if err != nil {
		return nil, err
	}

	per := newInfo.(*entity.Permission)
	var result pb.Permission
	err = copier.Copy(&result, per)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
