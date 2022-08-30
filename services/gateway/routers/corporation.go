package routers

import (
	rest "avatar/services/gateway/handler/rest"
	pb "avatar/services/gateway/protos/corporation"

	"github.com/gin-gonic/gin"
)

func InitCorporationsRouter(r gin.IRouter, corporationClient pb.CorporationClient) {
	r.GET("/:id", rest.GetCorporation(corporationClient))
	r.GET("/", rest.GetListCorporation(corporationClient))
	r.DELETE("/:id", rest.DeleteCorporation(corporationClient))
	r.POST("/", rest.CreateCorporation(corporationClient))
	r.PUT("/:id", rest.UpdateCorporation(corporationClient))
}
