package account_role

import (
	pb "avatar/services/gateway/protos/account_management"
	"context"
)

type DeleteRoleInput struct {
	Id int64 `json:"id" binding:"required"`
}

func DeleteRole(input *DeleteRoleInput, client pb.AccountManagementClient) (err error) {
	ctx := context.Background()
	_, err = client.DeleteAccountRole(ctx, &pb.DeleteAccountRoleRequest{
		Id: input.Id,
	})
	return err
}
