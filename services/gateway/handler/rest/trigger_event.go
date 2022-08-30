package rest

import (
	"avatar/services/gateway/config"
	"avatar/services/gateway/domain/broker"
	pb "avatar/services/gateway/protos/pos"
	"avatar/services/gateway/usecase/trigger_event"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TriggerEventPOSSide(client pb.POSClient, cfg *config.Environment, broker broker.Broker) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input trigger_event.TriggerEventPOSSideInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		err := trigger_event.TriggerEventPOSSide(&input, client, cfg, broker)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	}
}

func TriggerEventOperatorSide(client pb.POSClient, broker broker.Broker) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input trigger_event.TriggerEventOperatorSideInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := trigger_event.TriggerEventOperatorSide(&input, client, broker)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}
