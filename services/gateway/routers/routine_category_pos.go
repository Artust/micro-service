package routers

import (
	rest "avatar/services/gateway/handler/rest"
	pbPos "avatar/services/gateway/protos/pos"

	"github.com/gin-gonic/gin"
)

func InitRoutineCategoryPosRouter(r gin.IRouter, posClient pbPos.POSClient) {
	r.GET("/", rest.GetListRoutineCategoryPOS(posClient))
	r.GET("/:id", rest.GetRoutineCategoryPOS(posClient))
	r.POST("/", rest.CreateRoutineCategoryPOS(posClient))
	r.PUT("/:id", rest.UpdateRoutineCategoryPOS(posClient))
	r.DELETE("/:id", rest.DeleteRoutineCategoryPOS(posClient))
}
