package account

import (
	"avatar/services/gateway/config"
	pb "avatar/services/gateway/protos/account_management"
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/copier"
)

type UpdateAccountInput struct {
	Id       int64
	Email    string `json:"email"`
	Username string `json:"username"`
	Gender   int64  `json:"gender"`
	RoleId   int64  `json:"roleId"`
	CenterId int64  `json:"centerId"`
	Avatar   string `json:"avatar"`
}

func UpdateAccount(
	input *UpdateAccountInput,
	client pb.AccountManagementClient,
	cfg *config.Environment,
) (*Account, error) {
	ctx := context.Background()

	data := &pb.UpdateAccountRequest{}
	err := copier.Copy(&data, input)
	if err != nil {
		return nil, err
	}
	data.Avatar = data.Avatar[strings.LastIndex(input.Avatar, "/")+1:]
	acc, err := client.UpdateAccount(ctx, data)
	if err != nil {
		return nil, err
	}
	return &Account{
		Id:        acc.Id,
		Email:     acc.Email,
		Username:  acc.Username,
		Gender:    acc.Gender,
		RoleId:    acc.RoleId,
		CenterId:  acc.CenterId,
		Status:    acc.Status,
		CreatedBy: acc.CreatedBy,
		Avatar:    fmt.Sprintf("%s/%s/%s", cfg.S3Uri, config.AvatarAccountBucketName, acc.Avatar),
		CreatedAt: acc.CreatedAt,
		UpdatedAt: acc.UpdatedAt,
	}, err
}
