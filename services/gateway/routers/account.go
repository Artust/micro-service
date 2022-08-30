package routers

import (
	"avatar/services/gateway/config"
	rest "avatar/services/gateway/handler/rest"
	upload "avatar/services/gateway/infra/upload/respository"
	pb "avatar/services/gateway/protos/account_management"

	"github.com/gin-gonic/gin"
)

func InitAccountRouter(
	r gin.IRouter,
	accountManagementClient pb.AccountManagementClient,
	upload upload.UploadClient,
	cfg *config.Environment,
) {
	r.GET("/", rest.GetListAccount(accountManagementClient, cfg))
	r.GET("/:id", rest.GetAccount(accountManagementClient, cfg))
	r.POST("/", rest.CreateAccount(accountManagementClient, cfg))
	r.PUT("/:id", rest.UpdateAccount(accountManagementClient, cfg))
	r.PUT("/:id/change-password", rest.ChangePassword(accountManagementClient))
	r.PATCH("/:id/activation", rest.ActiveAccount(accountManagementClient))
	r.DELETE("/:id/activation", rest.DeactiveAccount(accountManagementClient))
}
