package permission

import (
	pb "avatar/services/gateway/protos/account_management"
	"context"
)

type DeletePermissionInput struct {
	Id int64 `json:"id" binding:"required"`
}

func DeletePermission(input *DeletePermissionInput, client pb.AccountManagementClient) (err error) {
	ctx := context.Background()
	_, err = client.DeletePermission(ctx, &pb.DeletePermissionRequest{
		Id: input.Id,
	})
	return err
}
