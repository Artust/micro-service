package routers

import (
	"avatar/services/gateway/config"
	rest "avatar/services/gateway/handler/rest"
	upload "avatar/services/gateway/infra/upload/respository"
	pbCenter "avatar/services/gateway/protos/center"

	"github.com/gin-gonic/gin"
)

func InitAvatarRouter(
	r gin.IRouter,
	centerClient pbCenter.CenterClient,
	upload upload.UploadClient,
	cfg *config.Environment,
) {
	r.PUT("/:id", rest.UpdateAvatar(centerClient, upload, cfg))
	r.GET("/", rest.GetListAvatar(centerClient, cfg))
	r.DELETE("/:id", rest.DeleteAvatar(centerClient))
	r.GET("/:id", rest.GetAvatar(centerClient, cfg))
	r.POST("/", rest.CreateAvatar(centerClient, upload, cfg))
}
