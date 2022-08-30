package rest

import (
	"avatar/services/gateway/domain/broker"
	pb "avatar/services/gateway/protos/talk_session"
	"avatar/services/gateway/usecase/note"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateNote(client pb.TalkSessionClient, broker broker.Broker) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input note.CreateNoteInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := note.CreateNote(&input, client, broker)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetListNote(client pb.TalkSessionClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input note.GetListNoteInput
		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := note.GetListNote(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func UpdateNote(client pb.TalkSessionClient, broker broker.Broker) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input note.UpdateNoteInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = id
		output, err := note.UpdateNote(&input, client, broker)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func DeleteNote(client pb.TalkSessionClient, broker broker.Broker) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input note.DeleteNoteInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = id
		output, err := note.DeleteNote(&input, client, broker)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}
