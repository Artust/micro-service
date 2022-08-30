package user_activity

import (
	pb "avatar/services/gateway/protos/account_management"
	"context"
)

type UserActivity struct {
	Id          int64  `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	AccountId   int64  `json:"accountId" binding:"required"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	DeletedAt   string `json:"deletedAt"`
}

func SaveUserActivity(input *UserActivity, client pb.AccountManagementClient) (*UserActivity, error) {
	ctx := context.Background()
	registerResponse, err := client.SaveUserActivity(ctx, &pb.UserActivity{
		Id:          input.Id,
		Name:        input.Name,
		Description: input.Description,
		AccountId:   input.AccountId,
	})
	if err != nil {
		return nil, err
	}
	return &UserActivity{
		Id:          input.Id,
		Name:        input.Name,
		Description: input.Description,
		AccountId:   input.AccountId,
		CreatedAt:   registerResponse.CreatedAt,
		UpdatedAt:   registerResponse.UpdatedAt,
	}, nil
}
