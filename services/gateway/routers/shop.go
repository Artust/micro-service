package routers

import (
	rest "avatar/services/gateway/handler/rest"
	pb "avatar/services/gateway/protos/corporation"

	"github.com/gin-gonic/gin"
)

func InitShopsRouter(r gin.IRouter, corporationClient pb.CorporationClient) {
	r.GET("/:id", rest.GetShop(corporationClient))
	r.GET("/", rest.GetListShop(corporationClient))
	r.DELETE("/:id", rest.DeleteShop(corporationClient))
	r.POST("/", rest.CreateShop(corporationClient))
	r.PUT("/:id", rest.UpdateShop(corporationClient))
}
