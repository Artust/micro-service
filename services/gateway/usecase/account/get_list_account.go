package account

import (
	"avatar/services/gateway/config"
	pb "avatar/services/gateway/protos/account_management"
	"context"
	"fmt"
)

type GetListAccountInput struct {
	Page     int64 `form:"page"`
	PerPage  int64 `form:"perPage"`
	Gender   int64 `form:"gender"`
	RoleId   int64 `form:"roleId"`
	CenterId int64 `form:"centerId"`
	Status   int64 `form:"status"`
}

type GetListAccountOuput struct {
	Accounts []*Account `json:"accounts"`
}

func GetListAccount(
	input *GetListAccountInput,
	client pb.AccountManagementClient,
	cfg *config.Environment,
) (*GetListAccountOuput, error) {
	ctx := context.Background()
	data := &pb.GetListAccountRequest{
		Page:     input.Page,
		PerPage:  input.PerPage,
		Gender:   input.Gender,
		RoleId:   input.RoleId,
		CenterId: input.CenterId,
		Status:   input.Status,
	}
	response, err := client.GetListAccount(ctx, data)
	if err != nil {
		return nil, err
	}
	var output GetListAccountOuput
	output.Accounts = make([]*Account, 0)
	for _, v := range response.Accounts {
		output.Accounts = append(output.Accounts, &Account{
			Id:        v.Id,
			Email:     v.Email,
			Username:  v.Username,
			Gender:    v.Gender,
			RoleId:    v.RoleId,
			CenterId:  v.CenterId,
			Status:    v.Status,
			Avatar:    fmt.Sprintf("%s/%s/%s", cfg.S3Uri, config.AvatarAccountBucketName, v.Avatar),
			CreatedBy: v.CreatedBy,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}
	return &output, nil
}
