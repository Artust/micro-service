package rest

import (
	"avatar/services/gateway/usecase/center"
	"net/http"
	"strconv"

	pb "avatar/services/gateway/protos/corporation"

	"github.com/gin-gonic/gin"
)

func CreateCenter(corporationClient pb.CorporationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input center.CreateCenterInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := center.CreateCenter(&input, corporationClient)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func UpdateCenter(corporationClient pb.CorporationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input center.UpdateCenterInput
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
		output, err := center.UpdateCenter(&input, corporationClient)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func DeleteCenter(corporationClient pb.CorporationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input center.DeleteCenterInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.ID = id
		output, err := center.DeleteCenter(&input, corporationClient)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetCenter(corporationClient pb.CorporationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input center.GetCenterInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.ID = id
		output, err := center.GetCenter(&input, corporationClient)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetListCenter(corporationClient pb.CorporationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input center.GetListCenterInput
		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := center.GetListCenter(&input, corporationClient)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}
