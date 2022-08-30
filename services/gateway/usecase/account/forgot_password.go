package account

import (
	pb "avatar/services/gateway/protos/account_management"
	"context"
)

type ForgotPasswordInput struct {
	Email string `json:"email" binding:"required"`
}

func ForgotPassword(input *ForgotPasswordInput, client pb.AccountManagementClient) (err error) {
	ctx := context.Background()
	_, err = client.ForgotPassword(ctx, &pb.ForgotPasswordRequest{
		Email: input.Email,
	})
	return err
}
