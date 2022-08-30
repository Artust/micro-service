package account

import (
	"avatar/services/gateway/config"
	pb "avatar/services/gateway/protos/account_management"
	"context"
	"fmt"

	"github.com/jinzhu/copier"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginOutput struct {
	Token       string `json:"token"`
	UserId      int64  `json:"userId"`
	DisplayName string `json:"displayName"`
	Avatar      string `json:"avatar"`
	Gender      int64  `json:"gender"`
}

func Login(
	input *LoginInput,
	client pb.AccountManagementClient,
	cfg *config.Environment) (*LoginOutput, error) {
	ctx := context.Background()
	login, err := client.Login(ctx, &pb.LoginRequest{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}
	var output LoginOutput
	err = copier.Copy(&output, login)
	if err != nil {
		return nil, err
	}
	output.Avatar = fmt.Sprintf("%s%s", cfg.S3Uri, output.Avatar)
	return &output, err
}
