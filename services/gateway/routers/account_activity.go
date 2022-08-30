package routers

import (
	rest "avatar/services/gateway/handler/rest"
	pb "avatar/services/gateway/protos/account_management"

	"github.com/gin-gonic/gin"
)

func InitAccountActivityRouter(r gin.IRouter, accountManagementClient pb.AccountManagementClient) {
	r.GET("/:id", rest.GetUserActivity(accountManagementClient))
	r.POST("/", rest.SaveUserActivity(accountManagementClient))
}
