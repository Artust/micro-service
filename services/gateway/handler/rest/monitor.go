package rest

import (
	"avatar/services/gateway/config"
	pb "avatar/services/gateway/protos/pos"
	"avatar/services/gateway/usecase/monitor"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateMonitor(client pb.POSClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input monitor.CreateMonitorInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := monitor.CreateMonitor(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetMonitor(client pb.POSClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input monitor.GetMonitorInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = id
		output, err := monitor.GetMonitor(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func UpdateMonitor(client pb.POSClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input monitor.UpdateMonitorInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = id
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := monitor.UpdateMonitor(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func DeleteMonitor(client pb.POSClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input monitor.DeleteMonitorInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = id
		output, err := monitor.DeleteMonitor(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetListMonitor(client pb.POSClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input monitor.GetListMonitorInput
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
		output, err := monitor.GetListMonitor(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}
