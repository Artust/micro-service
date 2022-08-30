package routers

import (
	"avatar/services/gateway/domain/broker"
	rest "avatar/services/gateway/handler/rest"
	pb "avatar/services/gateway/protos/talk_session"

	"github.com/gin-gonic/gin"
)

func InitNotesRouter(r gin.IRouter, client pb.TalkSessionClient, broker broker.Broker) {
	r.POST("/", rest.CreateNote(client, broker))
	r.GET("/", rest.GetListNote(client))
	r.PUT("/:id", rest.UpdateNote(client, broker))
	r.DELETE("/:id", rest.DeleteNote(client, broker))
}
