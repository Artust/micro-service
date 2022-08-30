package account

import (
	pb "avatar/services/gateway/protos/account_management"
	"context"

	"github.com/jinzhu/copier"
)

type ChangePasswordInput struct {
	Id          int64  `json:"id" binding:"required"`
	Password    string `json:"password" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

func ChangePassword(input *ChangePasswordInput, client pb.AccountManagementClient) (err error) {
	ctx := context.Background()

	data := &pb.ChangePasswordRequest{}
	err = copier.Copy(&data, input)
	if err != nil {
		return err
	}
	_, err = client.ChangePassword(ctx, data)
	return err
}
