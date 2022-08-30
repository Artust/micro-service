package account

import (
	pb "avatar/services/gateway/protos/account_management"
	"context"

	"google.golang.org/grpc/metadata"
)

type ResetPasswordInput struct {
	NewPassword        string `json:"newPassword" binding:"required"`
	ResetPasswordToken string `form:"token"`
}

type ResetPasswordOutput struct {
	Message string `json:"message"`
}

func ResetPassword(input *ResetPasswordInput, client pb.AccountManagementClient) (err error) {
	ctx := metadata.AppendToOutgoingContext(context.Background(), "reset_password_token", input.ResetPasswordToken)
	_, err = client.ResetPassword(ctx, &pb.ResetPasswordRequest{
		NewPassword: input.NewPassword,
	})
	return err
}
