package account_role

import (
	pb "avatar/services/gateway/protos/account_management"
	"context"
)

type CreateRoleInput struct {
	Name          string  `json:"name" binding:"required"`
	PermissionIds []int64 `json:"permissionIds" binding:"required"`
	Level         int64   `json:"level" binding:"required"`
}
type CreateRoleOutput struct {
	Id            int64   `json:"id"`
	Name          string  `json:"name"`
	PermissionIds []int64 `json:"permissionIds"`
	Level         int64   `json:"level"`
	CreatedAt     string  `json:"createdAt"`
	UpdatedAt     string  `json:"updatedAt"`
}

func CreateRole(input *CreateRoleInput, client pb.AccountManagementClient) (*CreateRoleOutput, error) {
	ctx := context.Background()
	data := &pb.CreateAccountRoleRequest{
		Name:          input.Name,
		Level:         input.Level,
		PermissionIds: input.PermissionIds,
	}
	registerResponse, err := client.CreateAccountRole(ctx, data)
	if err != nil {
		return nil, err
	}
	return &CreateRoleOutput{
		Id:            registerResponse.Id,
		Name:          registerResponse.Name,
		PermissionIds: registerResponse.PermissionIds,
		Level:         registerResponse.Level,
		CreatedAt:     registerResponse.CreatedAt,
		UpdatedAt:     registerResponse.UpdatedAt,
	}, nil
}
