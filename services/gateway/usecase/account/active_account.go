package account

import (
	pb "avatar/services/gateway/protos/account_management"
	"context"
)

type ActiveAccountInput struct {
	Id int64 `json:"id" binding:"required"`
}

func ActiveAccount(input *ActiveAccountInput, client pb.AccountManagementClient) (err error) {
	ctx := context.Background()
	_, err = client.ActiveAccount(ctx, &pb.ActiveAccountRequest{
		Id: input.Id,
	})
	return err
}
