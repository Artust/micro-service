package routers

import (
	"avatar/services/gateway/config"
	rest "avatar/services/gateway/handler/rest"
	upload "avatar/services/gateway/infra/upload/respository"
	pb "avatar/services/gateway/protos/pos"

	"github.com/gin-gonic/gin"
)

func InitRoutinesRouter(
	r gin.IRouter,
	posClient pb.POSClient,
	upload upload.UploadClient,
	cfg *config.Environment,
) {
	r.PUT("/:id", rest.UpdateRoutinePos(posClient))
	r.GET("/", rest.GetListRoutinePos(posClient, cfg))
	r.DELETE("/:id", rest.DeleteRoutinePos(posClient))
	r.GET("/:id", rest.GetRoutinePos(posClient, cfg))
	r.POST("/", rest.CreateRoutinePos(posClient, upload, cfg))

	InitRoutineCategoryPosRouter(r.Group("categories"), posClient)
}
