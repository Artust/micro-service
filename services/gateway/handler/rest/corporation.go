package rest

import (
	"avatar/services/gateway/config"
	"avatar/services/gateway/usecase/corporation"
	"net/http"
	"strconv"

	pb "avatar/services/gateway/protos/corporation"

	"github.com/gin-gonic/gin"
)

func CreateCorporation(client pb.CorporationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input corporation.CreateCorporationInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := corporation.CreateCorporation(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetCorporation(client pb.CorporationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input corporation.GetCorporationInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = id
		output, err := corporation.GetCorporation(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetListCorporation(client pb.CorporationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input corporation.GetListCorporationInput
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
		output, err := corporation.GetListCorporation(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func UpdateCorporation(client pb.CorporationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input corporation.UpdateCorporationInput
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
		output, err := corporation.UpdateCorporation(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func DeleteCorporation(client pb.CorporationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input corporation.DeleteCorporationInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = id
		output, err := corporation.DeleteCorporation(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}
