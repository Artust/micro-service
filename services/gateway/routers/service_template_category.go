package routers

import (
	rest "avatar/services/gateway/handler/rest"
	pbCenter "avatar/services/gateway/protos/center"

	"github.com/gin-gonic/gin"
)

func InitServiceTemplateCategoryRouter(
	r gin.IRouter,
	centerClient pbCenter.CenterClient,
) {
	r.PUT("/:id", rest.UpdateServiceTemplateCategory(centerClient))
	r.GET("/", rest.GetListServiceTemplateCategory(centerClient))
	r.DELETE("/:id", rest.DeleteServiceTemplateCategory(centerClient))
	r.GET("/:id", rest.GetServiceTemplateCategory(centerClient))
	r.POST("/", rest.CreateServiceTemplateCategory(centerClient))
}
