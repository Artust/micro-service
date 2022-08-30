package permission

import (
	"avatar/services/account_management/config"
	"avatar/services/account_management/domain/repository"
	pb "avatar/services/account_management/protos"
	"context"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func Delete(
	ctx context.Context,
	db neo4j.Driver,
	permissionRepository repository.PermissionRepository,
	input *pb.DeletePermissionRequest,
) (*pb.Empty, error) {
	session := db.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		ctx = context.WithValue(ctx, config.Neo4jTransactionContextKey, tx)
		err := permissionRepository.Delete(ctx, input.Id)
		return nil, err
	})
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
