package routers

import (
	rest "avatar/services/gateway/handler/rest"
	pbCenter "avatar/services/gateway/protos/center"

	"github.com/gin-gonic/gin"
)

func InitRoutineCategoryServiceTemplateRouter(
	r gin.IRouter,
	centerClient pbCenter.CenterClient,
) {
	r.GET("/", rest.GetListRoutineCategory(centerClient))
	r.GET("/:id", rest.GetRoutineCategory(centerClient))
	r.POST("/", rest.CreateRoutineCategory(centerClient))
	r.PUT("/:id", rest.UpdateRoutineCategory(centerClient))
	r.DELETE("/:id", rest.DeleteRoutineCategory(centerClient))
}
