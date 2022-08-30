package routers

import (
	"avatar/services/gateway/handler/rest"
	upload "avatar/services/gateway/infra/upload/respository"

	"github.com/gin-gonic/gin"
)

func InitUploadFile(r gin.IRouter, upload upload.UploadClient) {
	r.POST("/:bucket", rest.UploadFile(upload))
}
