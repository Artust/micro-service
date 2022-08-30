package routers

import (
	"avatar/services/gateway/config"
	"avatar/services/gateway/domain/broker"
	rest "avatar/services/gateway/handler/rest"
	pb "avatar/services/gateway/protos/pos"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
)

func InitTriggerEventRouter(r gin.IRouter,
	posClient pb.POSClient,
	s3Session *session.Session,
	cfg *config.Environment,
	broker broker.Broker,
	) {
	r.POST("/trigger-event-pos-side", rest.TriggerEventPOSSide(posClient, cfg, broker))
	r.POST("/trigger-event-operator-side", rest.TriggerEventOperatorSide(posClient, broker))
}
