package routers

import (
	"avatar/services/gateway/config"
	"avatar/services/gateway/middleware"
	pbAccountManagement "avatar/services/gateway/protos/account_management"
	pbCenter "avatar/services/gateway/protos/center"
	pbCorporation "avatar/services/gateway/protos/corporation"
	pbPos "avatar/services/gateway/protos/pos"
	pbStreaming "avatar/services/gateway/protos/streaming"
	pbTalkSession "avatar/services/gateway/protos/talk_session"
	"fmt"
	"strconv"
	"strings"
	"time"

	"avatar/services/gateway/domain/broker"
	upload "avatar/services/gateway/infra/upload/respository"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RestServer struct {
	StreamingClient         pbStreaming.StreamingClient
	AccountManagementClient pbAccountManagement.AccountManagementClient
	PosClient               pbPos.POSClient
	CenterClient            pbCenter.CenterClient
	TalkSessionClient       pbTalkSession.TalkSessionClient
	CorporationClient       pbCorporation.CorporationClient
	BrokerClient            broker.Broker
	Config                  *config.Environment
	Engine                  *gin.Engine
	S3Session               *session.Session
	Upload                  upload.UploadClient
}

func InitRouter(restServer *RestServer) {
	restServer.Engine.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Split(restServer.Config.AllowOrigins, ","),
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Access-Control-Allow-Headers", "Authorization", "X-XSRF-TOKEN"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	api := restServer.Engine.Group("/api")
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	api.Use(middleware.AuthorizeJWT())
	{
		InitPosRouter(
			api.Group("pos"),
			restServer.PosClient,
			restServer.CenterClient,
			restServer.S3Session,
			restServer.Config,
			restServer.BrokerClient,
			restServer.Upload,
		)
		InitNotesRouter(
			api.Group("talk-session-notes"),
			restServer.TalkSessionClient,
			restServer.BrokerClient,
		)
		InitTriggerEventRouter(
			api,
			restServer.PosClient,
			restServer.S3Session,
			restServer.Config,
			restServer.BrokerClient,
		)
		InitShopsRouter(api.Group("shops"), restServer.CorporationClient)
		InitDevicesRouter(api.Group("devices"), restServer.CorporationClient)
		InitCenterRouter(api.Group("centers"), restServer.CorporationClient)
		InitCorporationsRouter(api.Group("corporations"), restServer.CorporationClient)
		InitServiceTemplateRouter(
			api.Group("service-templates"),
			restServer.CenterClient,
			restServer.Upload,
			restServer.Config,
		)
		InitAvatarRouter(
			api.Group("avatars"),
			restServer.CenterClient,
			restServer.Upload,
			restServer.Config,
		)
		InitMonitorsRouter(api.Group("monitors"), restServer.PosClient)
		InitAuthenticateRouter(
			api.Group("auth"),
			restServer.AccountManagementClient,
			restServer.Config,
		)
		InitAccountRouter(
			api.Group("account"),
			restServer.AccountManagementClient,
			restServer.Upload,
			restServer.Config,
		)
		InitPermissionRouter(api.Group("permission"), restServer.AccountManagementClient)
		InitAccountRoleRouter(api.Group("account-role"), restServer.AccountManagementClient)
		InitAccountActivityRouter(api.Group("account/activity"), restServer.AccountManagementClient)
		InitIpCamerasRouter(api.Group("ip-cameras"), restServer.PosClient, restServer.CorporationClient)
		InitRtspStream(api.Group("rtsp-stream"), restServer.PosClient, restServer.CorporationClient, restServer.Config)
		InitUploadFile(api.Group("upload"), restServer.Upload)
	}
	restServer.Engine.Run(fmt.Sprintf(":%v", strconv.Itoa(restServer.Config.GatewayRestPort)))
}
