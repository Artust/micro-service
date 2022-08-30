package routers

import (
	"avatar/services/gateway/config"
	"avatar/services/gateway/domain/broker"
	rest "avatar/services/gateway/handler/rest"
	upload "avatar/services/gateway/infra/upload/respository"
	pbCenter "avatar/services/gateway/protos/center"
	pbPos "avatar/services/gateway/protos/pos"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
)

func InitPosRouter(
	r gin.IRouter,
	posClient pbPos.POSClient,
	centerClient pbCenter.CenterClient,
	s3Session *session.Session,
	cfg *config.Environment,
	broker broker.Broker,
	upload upload.UploadClient,
) {

	r.GET("/:id/routines", rest.GetListRoutineByCategory(posClient, cfg))
	r.GET("/", rest.GetListPos(posClient, cfg))
	r.GET("/:id", rest.GetPos(posClient, cfg))
	r.POST("/", rest.CreatePos(posClient, centerClient, cfg))
	r.PUT("/:id", rest.UpdatePos(posClient, centerClient, cfg))
	r.DELETE("/:id", rest.DeletePos(posClient))

	InitRoutinesRouter(r.Group("routines"), posClient, upload, cfg)
}
