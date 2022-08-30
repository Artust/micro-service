package rest

import (
	pb "avatar/services/gateway/protos/corporation"
	"avatar/services/gateway/usecase/device"
	"avatar/services/gateway/config"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateDevice(client pb.CorporationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input device.CreateDeviceInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := device.CreateDevice(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetDevice(client pb.CorporationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input device.GetDeviceInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = id
		output, err := device.GetDevice(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func UpdateDevice(client pb.CorporationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input device.UpdateDeviceInput
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
		output, err := device.UpdateDevice(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func DeleteDevice(client pb.CorporationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input device.DeleteDeviceInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = id
		output, err := device.DeleteDevice(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetListDevice(client pb.CorporationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input device.GetListDeviceInput
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
		output, err := device.GetListDevice(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}
