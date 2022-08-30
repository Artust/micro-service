package account

import (
	"avatar/services/gateway/config"
	pb "avatar/services/gateway/protos/account_management"
	"context"
	"fmt"
)

type GetAccountInput struct {
	Id int64 `form:"id" binding:"required"`
}

type Account struct {
	Id        int64  `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Gender    int64  `json:"gender"`
	RoleId    int64  `json:"roleId"`
	CenterId  int64  `json:"centerId"`
	Status    int64  `json:"status"`
	CreatedBy int64  `json:"createdBy"`
	Avatar    string `json:"avatar"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func GetAccount(
	input *GetAccountInput,
	client pb.AccountManagementClient,
	cfg *config.Environment,
) (*Account, error) {
	ctx := context.Background()
	acc, err := client.GetAccount(ctx, &pb.Id{
		Id: input.Id,
	})
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
