package routers

import (
	"avatar/services/gateway/config"
	rest "avatar/services/gateway/handler/rest"
	upload "avatar/services/gateway/infra/upload/respository"
	pbCenter "avatar/services/gateway/protos/center"

	"github.com/gin-gonic/gin"
)

func InitServiceTemplateRouter(
	r gin.IRouter,
	centerClient pbCenter.CenterClient,
	upload upload.UploadClient,
	cfg *config.Environment,
) {
	InitRoutineServiceTemplateRouter(r.Group("routines"), centerClient, upload, cfg)
	InitRoutineCategoryServiceTemplateRouter(r.Group("routines-categories"), centerClient)
	InitServiceTemplateCategoryRouter(r.Group("categories"), centerClient)
	r.PUT("/:id", rest.UpdateServiceTemplate(centerClient, cfg))
	r.GET("/", rest.GetListServiceTemplate(centerClient, cfg))
	r.DELETE("/:id", rest.DeleteServiceTemplate(centerClient))
	r.GET("/:id", rest.GetServiceTemplate(centerClient, cfg))
	r.POST("/", rest.CreateServiceTemplate(centerClient, upload, cfg))
	r.GET("/list-routines", rest.GetListRoutineCenterByCategory(centerClient, cfg))
}
