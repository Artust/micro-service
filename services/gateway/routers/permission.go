package routers

import (
	rest "avatar/services/gateway/handler/rest"
	pb "avatar/services/gateway/protos/account_management"

	"github.com/gin-gonic/gin"
)

func InitPermissionRouter(r gin.IRouter, accountManagementClient pb.AccountManagementClient) {
	r.GET("/", rest.GetListPermission(accountManagementClient))
	r.GET("/:id", rest.GetPermission(accountManagementClient))
	r.POST("/", rest.CreatePermission(accountManagementClient))
	r.PUT("/:id", rest.UpdatePermission(accountManagementClient))
	r.DELETE("/:id", rest.DeletePermission(accountManagementClient))
}
