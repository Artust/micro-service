package permission

import (
	pb "avatar/services/gateway/protos/account_management"
	"context"
)

type CreatePermissionInput struct {
	Entity           string `json:"entity" binding:"required"`
	PermissionAction string `json:"permissionAction" binding:"required"`
}
type CreatePermissionOutput struct {
	Id               int64  `json:"id"`
	Entity           string `json:"entity"`
	PermissionAction string `json:"permissionAction"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
}

func CreatePermission(input *CreatePermissionInput, client pb.AccountManagementClient) (*CreatePermissionOutput, error) {
	ctx := context.Background()
	registerResponse, err := client.CreatePermission(ctx, &pb.CreatePermissionRequest{
		Entity:           input.Entity,
		PermissionAction: input.PermissionAction,
	})
	if err != nil {
		return nil, err
	}
	return &CreatePermissionOutput{
		Id:               registerResponse.Id,
		Entity:           registerResponse.Entity,
		PermissionAction: registerResponse.PermissionAction,
		CreatedAt:        registerResponse.CreatedAt,
		UpdatedAt:        registerResponse.UpdatedAt,
	}, nil
}
