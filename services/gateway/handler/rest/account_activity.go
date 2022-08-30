package rest

import (
	pb "avatar/services/gateway/protos/account_management"
	"strconv"

	"avatar/services/gateway/usecase/user_activity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserActivity(client pb.AccountManagementClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input user_activity.GetUserActivityInput

		id := c.Param("id")
		newId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.AccountId = newId

		output, err := user_activity.GetUserActivity(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func SaveUserActivity(client pb.AccountManagementClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input user_activity.UserActivity
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := user_activity.SaveUserActivity(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}
