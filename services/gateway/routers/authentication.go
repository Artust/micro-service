package routers

import (
	"avatar/services/gateway/config"
	rest "avatar/services/gateway/handler/rest"
	pb "avatar/services/gateway/protos/account_management"

	"github.com/gin-gonic/gin"
)

func InitAuthenticateRouter(
	r gin.IRouter,
	accountManagementClient pb.AccountManagementClient,
	cfg *config.Environment,
) {
	r.POST("/login", rest.Login(accountManagementClient, cfg))
	r.POST("/forgot-password", rest.ForgotPassword(accountManagementClient))
	r.POST("/reset-password", rest.ResetPassword(accountManagementClient))
}
