package account_role

import (
	pb "avatar/services/gateway/protos/account_management"
	"context"
)

type GetRoleInput struct {
	Id int64 `json:"id" binding:"required"`
}
type AccountRole struct {
	Id            int64   `json:"id"`
	Name          string  `json:"name"`
	PermissionIds []int64 `json:"permissionIds"`
	Level         int64   `json:"level"`
	CreatedAt     string  `json:"createdAt"`
	UpdatedAt     string  `json:"updatedAt"`
	DeletedAt     string  `json:"deletedAt"`
}

func GetRole(input *GetRoleInput, client pb.AccountManagementClient) (*AccountRole, error) {
	ctx := context.Background()
	registerResponse, err := client.GetAccountRole(ctx, &pb.Id{
		Id: input.Id,
	})
	if err != nil {
		return nil, err
	}
	return &AccountRole{
		Id:            registerResponse.Id,
		Name:          registerResponse.Name,
		PermissionIds: registerResponse.PermissionIds,
		Level:         registerResponse.Level,
		CreatedAt:     registerResponse.CreatedAt,
		UpdatedAt:     registerResponse.UpdatedAt,
	}, nil
}
