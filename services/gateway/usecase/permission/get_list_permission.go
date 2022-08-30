package permission

import (
	pb "avatar/services/gateway/protos/account_management"
	"context"
)

type GetListPermissionOutout struct {
	Permissions []*Permission `json:"Permissions"`
}

func GetListPermission(client pb.AccountManagementClient) (*GetListPermissionOutout, error) {
	ctx := context.Background()
	response, err := client.GetListPermission(ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}
	var output GetListPermissionOutout
	output.Permissions = make([]*Permission, 0)
	for _, v := range response.Permissions {
		output.Permissions = append(output.Permissions, &Permission{
			Id:               v.Id,
			Entity:           v.Entity,
			PermissionAction: v.PermissionAction,
			CreatedAt:        v.CreatedAt,
			UpdatedAt:        v.UpdatedAt,
		})
	}
	return &output, nil
}
