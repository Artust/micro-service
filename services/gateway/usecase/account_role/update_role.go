package account_role

import (
	pb "avatar/services/gateway/protos/account_management"
	"context"
)

type UpdateRoleInput struct {
	Id            int64
	Name          string  `json:"name" binding:"required"`
	PermissionIds []int64 `json:"permissionIds" binding:"required"`
	Level         int64   `json:"level" binding:"required"`
}

func UpdateRole(input *UpdateRoleInput, client pb.AccountManagementClient) (*AccountRole, error) {
	ctx := context.Background()
	registerResponse, err := client.UpdateAccountRole(ctx, &pb.UpdateAccountRoleRequest{
		Id:            input.Id,
		Name:          input.Name,
		PermissionIds: input.PermissionIds,
		Level:         input.Level,
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
