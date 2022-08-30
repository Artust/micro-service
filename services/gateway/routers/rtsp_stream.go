package routers

import (
	"avatar/services/gateway/config"
	rest "avatar/services/gateway/handler/rest"
	pbCorporation "avatar/services/gateway/protos/corporation"
	pbPos "avatar/services/gateway/protos/pos"

	"github.com/gin-gonic/gin"
)

func InitRtspStream(
	r gin.IRouter,
	posClient pbPos.POSClient,
	corporationClient pbCorporation.CorporationClient,
	cfg *config.Environment,
) {
	r.GET("/", rest.GetRtspStream(posClient, corporationClient, cfg))
}
