package account_role

import (
	pb "avatar/services/gateway/protos/account_management"
	"context"
)

type GetListRoleOutput struct {
	AccountRoles []*AccountRole `json:"Roles"`
}

func GetListAccountRole(client pb.AccountManagementClient) (*GetListRoleOutput, error) {
	ctx := context.Background()
	response, err := client.GetListAccountRole(ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}
	var output GetListRoleOutput
	output.AccountRoles = make([]*AccountRole, 0)
	for _, v := range response.AccountRoles {
		output.AccountRoles = append(output.AccountRoles, &AccountRole{
			Id:            v.Id,
			Name:          v.Name,
			PermissionIds: v.PermissionIds,
			Level:         v.Level,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		})
	}
	return &output, nil
}
