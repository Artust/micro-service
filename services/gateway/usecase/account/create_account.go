package account

import (
	"avatar/pkg/jwt"
	"avatar/services/gateway/config"
	pb "avatar/services/gateway/protos/account_management"
	"context"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

type CreateAccountInput struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Gender   int64  `json:"gender" binding:"required"`
	RoleId   int64  `json:"roleId" binding:"required"`
	CenterId int64  `json:"centerId" binding:"required"`
	Avatar   string `json:"avatar" binding:"required"`
}

func CreateAccount(
	c *gin.Context,
	input *CreateAccountInput,
	client pb.AccountManagementClient,
	cfg *config.Environment,
) (*Account, error) {
	ctx := context.Background()
	claims, ok := c.Get("claims")
	if !ok {
		return nil, fmt.Errorf("user not authorized")
	}
	infoToken := claims.(*jwt.Claims)
	ctx = metadata.AppendToOutgoingContext(ctx, "uid", fmt.Sprint(infoToken.UserId))

	acc, err := client.CreateAccount(ctx, &pb.CreateAccountRequest{
		Email:    input.Email,
		Username: input.Username,
		RoleId:   input.RoleId,
		CenterId: input.CenterId,
		Gender:   input.Gender,
		Avatar:   input.Avatar[strings.LastIndex(input.Avatar, "/")+1:],
	})
	if err != nil {
		return nil, err
	}
	return &Account{
		Id:        acc.Id,
		Email:     acc.Email,
		Username:  acc.Username,
		Gender:    acc.Gender,
		RoleId:    acc.RoleId,
		CenterId:  acc.CenterId,
		Status:    acc.Status,
		CreatedBy: acc.CreatedBy,
		Avatar:    fmt.Sprintf("%s/%s/%s", cfg.S3Uri, config.AvatarAccountBucketName, acc.Avatar),
		CreatedAt: acc.CreatedAt,
		UpdatedAt: acc.UpdatedAt,
	}, nil
}
