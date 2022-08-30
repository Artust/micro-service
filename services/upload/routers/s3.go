package routers

import (
	"avatar/services/upload/config"
	"avatar/services/upload/handler"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
)

func UploadS3Router(r gin.IRouter, s3 *session.Session, env *config.Environment) {
	r.POST("/:bucket", handler.UploadToS3(s3, env))
}
