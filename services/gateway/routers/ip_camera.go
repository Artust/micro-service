package routers

import (
	rest "avatar/services/gateway/handler/rest"
	corporation "avatar/services/gateway/protos/corporation"
	pos "avatar/services/gateway/protos/pos"

	"github.com/gin-gonic/gin"
)

func InitIpCamerasRouter(
	r gin.IRouter,
	posClient pos.POSClient,
	corporationClient corporation.CorporationClient,
) {
	r.POST("/", rest.CreateIpCamera(posClient, corporationClient))
	r.GET("/:id", rest.GetIpCamera(posClient, corporationClient))
	r.PUT("/:id", rest.UpdateIpCamera(posClient, corporationClient))
	r.DELETE("/:id", rest.DeleteIpCamera(posClient, corporationClient))
	r.GET("/", rest.GetListIpCamera(posClient, corporationClient))
}
