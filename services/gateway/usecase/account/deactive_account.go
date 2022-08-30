package account

import (
	pb "avatar/services/gateway/protos/account_management"
	"context"
)

type DeactiveAccountInput struct {
	Id int64 `json:"id" binding:"required"`
}

func DeactiveAccount(input *DeactiveAccountInput, client pb.AccountManagementClient) (err error) {
	ctx := context.Background()
	_, err = client.DeactiveAccount(ctx, &pb.DeactiveAccountRequest{
		Id: input.Id,
	})
	return err
}
