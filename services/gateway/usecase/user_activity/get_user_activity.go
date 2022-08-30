package user_activity

import (
	pb "avatar/services/gateway/protos/account_management"
	"context"
)

type GetUserActivityInput struct {
	AccountId int64 `json:"accountId" binding:"required"`
}

type GetUserActivityOuput struct {
	Activities []*UserActivity `json:"activities"`
}

func GetUserActivity(input *GetUserActivityInput, client pb.AccountManagementClient) (*GetUserActivityOuput, error) {
	ctx := context.Background()
	response, err := client.GetUserActivity(ctx, &pb.AccountId{
		AccountId: input.AccountId,
	})
	if err != nil {
		return nil, err
	}
	var output GetUserActivityOuput
	output.Activities = make([]*UserActivity, 0)
	for _, v := range response.Activities {
		output.Activities = append(output.Activities, &UserActivity{
			Id:          v.Id,
			Name:        v.Name,
			Description: v.Description,
			AccountId:   v.AccountId,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}
	return &output, nil
}
