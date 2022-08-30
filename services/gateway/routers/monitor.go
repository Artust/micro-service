package routers

import (
	rest "avatar/services/gateway/handler/rest"
	pb "avatar/services/gateway/protos/pos"

	"github.com/gin-gonic/gin"
)

func InitMonitorsRouter(r gin.IRouter, posClient pb.POSClient) {
	r.POST("/", rest.CreateMonitor(posClient))
	r.GET("/:id", rest.GetMonitor(posClient))
	r.PUT("/:id", rest.UpdateMonitor(posClient))
	r.DELETE("/:id", rest.DeleteMonitor(posClient))
	r.GET("/", rest.GetListMonitor(posClient))
}
