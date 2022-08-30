package routers

import (
	"avatar/services/gateway/config"
	rest "avatar/services/gateway/handler/rest"
	upload "avatar/services/gateway/infra/upload/respository"
	pbCenter "avatar/services/gateway/protos/center"

	"github.com/gin-gonic/gin"
)

func InitRoutineServiceTemplateRouter(
	r gin.IRouter,
	centerClient pbCenter.CenterClient,
	upload upload.UploadClient,
	cfg *config.Environment,
) {
	r.GET("", rest.GetListRoutineCenter(centerClient, cfg))
	r.GET("/:id", rest.GetRoutineCenter(centerClient, cfg))
	r.POST("/", rest.CreateRoutineCenter(centerClient, upload, cfg))
	r.DELETE("/:id", rest.DeleteRoutineCenter(centerClient))
	r.PUT("/:id", rest.UpdateRoutineCenter(centerClient, cfg))
}
