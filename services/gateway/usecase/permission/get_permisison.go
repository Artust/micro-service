package permission

import (
	pb "avatar/services/gateway/protos/account_management"
	"context"
)

type GetPermissionInput struct {
	Id int64 `json:"id" binding:"required"`
}
type Permission struct {
	Id               int64  `json:"id"`
	Entity           string `json:"entity"`
	PermissionAction string `json:"permissionAction"`
	CreatedAt        string `json:"createdAt"`
	UpdatedAt        string `json:"updatedAt"`
	DeletedAt        string `json:"deletedAt"`
}

func GetPermission(input *GetPermissionInput, client pb.AccountManagementClient) (*Permission, error) {
	ctx := context.Background()
	registerResponse, err := client.GetPermission(ctx, &pb.Id{
		Id: input.Id,
	})
	if err != nil {
		return nil, err
	}
	return &Permission{
		Id:               registerResponse.Id,
		Entity:           registerResponse.Entity,
		PermissionAction: registerResponse.PermissionAction,
		CreatedAt:        registerResponse.CreatedAt,
		UpdatedAt:        registerResponse.UpdatedAt,
	}, nil
}
