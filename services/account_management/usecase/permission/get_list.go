package permission

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
	permissionRepository repository.PermissionRepository,
) (*pb.PermissionList, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	permissionsRaw, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		listPermission, errReset := permissionRepository.GetList(ctx)
		if errReset != nil {
			return nil, errReset
		}
		return listPermission, nil
	})
	if err != nil {
		return nil, err
	}
	permissionList := permissionsRaw.([]*entity.Permission)
	var results pb.PermissionList
	results.Permissions = make([]*pb.Permission, 0)
	for _, acc := range permissionList {
		var response pb.Permission
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
		results.Permissions = append(results.Permissions, &response)
	}
	return &results, nil
}
