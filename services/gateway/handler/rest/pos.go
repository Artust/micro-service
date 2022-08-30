package rest

import (
	"avatar/services/gateway/config"
	pbCenter "avatar/services/gateway/protos/center"
	pbPos "avatar/services/gateway/protos/pos"
	"avatar/services/gateway/usecase/pos"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePos(
	posClient pbPos.POSClient,
	centerClient pbCenter.CenterClient,
	cfg *config.Environment,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input pos.CreatePosInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := pos.CreatePos(&input, posClient, centerClient, cfg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetPos(posClient pbPos.POSClient, cfg *config.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input pos.GetPosInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = id
		output, err := pos.GetPos(&input, posClient, cfg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetListPos(posClient pbPos.POSClient, cfg *config.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input pos.GetListPosInput
		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if input.Page == 0 {
			input.Page = config.Page
		}
		if input.PerPage == 0 {
			input.PerPage = config.PerPage
		}
		output, err := pos.GetListPos(&input, posClient, cfg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func UpdatePos(
	posClient pbPos.POSClient,
	centerClient pbCenter.CenterClient,
	cfg *config.Environment,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input pos.UpdatePosInput
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
		output, err := pos.UpdatePos(&input, posClient, centerClient, cfg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func DeletePos(posClient pbPos.POSClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input pos.DeletePosInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = id
		output, err := pos.DeletePos(&input, posClient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}
