package rest

import (
	"avatar/services/gateway/config"
	corporation "avatar/services/gateway/protos/corporation"
	pos "avatar/services/gateway/protos/pos"
	"avatar/services/gateway/usecase/ip_camera"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateIpCamera(posClient pos.POSClient, corporationClient corporation.CorporationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input ip_camera.CreateIpCameraInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := ip_camera.CreateIpCamera(&input, posClient, corporationClient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetIpCamera(posClient pos.POSClient, corporationClient corporation.CorporationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input ip_camera.GetIpCameraInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = id
		output, err := ip_camera.GetIpCamera(&input, posClient, corporationClient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func UpdateIpCamera(posClient pos.POSClient, corporationClient corporation.CorporationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input ip_camera.UpdateIpCameraInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.CameraId = id
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := ip_camera.UpdateIpCamera(&input, posClient, corporationClient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func DeleteIpCamera(posClient pos.POSClient, corporationClient corporation.CorporationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input ip_camera.DeleteIpCameraInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = id
		output, err := ip_camera.DeleteIpCamera(&input, posClient, corporationClient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetListIpCamera(posClient pos.POSClient, corporationClient corporation.CorporationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input ip_camera.GetListIpCameraInput
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
		output, err := ip_camera.GetListIpCamera(&input, posClient, corporationClient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}
