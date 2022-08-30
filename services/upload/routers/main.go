package routers

import (
	"avatar/services/upload/config"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RestServer struct {
	Config    *config.Environment
	Engine    *gin.Engine
	S3Session *session.Session
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
	api.Use()
	{
		UploadS3Router(api.Group("s3"), restServer.S3Session, restServer.Config)
	}
	restServer.Engine.Run(fmt.Sprintf(":%v", restServer.Config.UploadPort))
}
