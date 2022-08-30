package routers

import (
	rest "avatar/services/gateway/handler/rest"
	pb "avatar/services/gateway/protos/account_management"

	"github.com/gin-gonic/gin"
)

func InitAccountRoleRouter(r gin.IRouter, accountManagementClient pb.AccountManagementClient) {
	r.GET("/", rest.GetListAccountRole(accountManagementClient))
	r.POST("/", rest.CreateAccountRole(accountManagementClient))
	r.GET("/:id", rest.GetAccountRole(accountManagementClient))
	r.PUT("/:id", rest.UpdateAccountRole(accountManagementClient))
	r.DELETE("/:id", rest.DeleteAccountRole(accountManagementClient))
}
