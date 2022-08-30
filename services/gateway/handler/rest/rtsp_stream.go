package rest

import (
	"avatar/services/gateway/config"
	corporation "avatar/services/gateway/protos/corporation"
	pos "avatar/services/gateway/protos/pos"
	"avatar/services/gateway/usecase/rtsp_stream"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRtspStream(
	posClient pos.POSClient,
	corporationClient corporation.CorporationClient,
	cfg *config.Environment,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		output, err := rtsp_stream.GetRtspStream(posClient, corporationClient, cfg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}
