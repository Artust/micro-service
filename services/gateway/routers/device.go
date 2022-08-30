package routers

import (
	rest "avatar/services/gateway/handler/rest"
	pb "avatar/services/gateway/protos/corporation"

	"github.com/gin-gonic/gin"
)

func InitDevicesRouter(r gin.IRouter, corporationClient pb.CorporationClient) {
	r.POST("/", rest.CreateDevice(corporationClient))
	r.GET("/:id", rest.GetDevice(corporationClient))
	r.PUT("/:id", rest.UpdateDevice(corporationClient))
	r.DELETE("/:id", rest.DeleteDevice(corporationClient))
	r.GET("/", rest.GetListDevice(corporationClient))
}
