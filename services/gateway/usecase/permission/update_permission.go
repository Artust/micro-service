package permission

import (
	pb "avatar/services/gateway/protos/account_management"
	"context"
)

type UpdatePermissionInput struct {
	Id               int64
	Entity           string `json:"entity" binding:"required"`
	PermissionAction string `json:"permissionAction" binding:"required"`
}
type UpdatePermissionOutput struct {
	Id               int64  `json:"id"`
	Entity           string `json:"entity"`
	PermissionAction string `json:"permissionAction"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
}

func UpdatePermission(input *UpdatePermissionInput, client pb.AccountManagementClient) (*UpdatePermissionOutput, error) {
	ctx := context.Background()
	registerResponse, err := client.UpdatePermission(ctx, &pb.UpdatePermissionRequest{
		Id:               input.Id,
		Entity:           input.Entity,
		PermissionAction: input.PermissionAction,
	})
	if err != nil {
		return nil, err
	}
	return &UpdatePermissionOutput{
		Id:               registerResponse.Id,
		Entity:           registerResponse.Entity,
		PermissionAction: registerResponse.PermissionAction,
		CreatedAt:        registerResponse.CreatedAt,
		UpdatedAt:        registerResponse.UpdatedAt,
	}, nil
}
