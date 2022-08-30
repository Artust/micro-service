package routers

import (
	"avatar/services/gateway/handler/rest"
	pb "avatar/services/gateway/protos/corporation"

	"github.com/gin-gonic/gin"
)

func InitCenterRouter(r gin.IRouter, corporationClient pb.CorporationClient) {
	r.POST("/", rest.CreateCenter(corporationClient))
	r.GET("/:id", rest.GetCenter(corporationClient))
	r.PUT("/:id", rest.UpdateCenter(corporationClient))
	r.DELETE("/:id", rest.DeleteCenter(corporationClient))
	r.GET("/", rest.GetListCenter(corporationClient))
}
