package rest

import (
	pb "avatar/services/gateway/protos/corporation"
	"avatar/services/gateway/usecase/shop"
	"avatar/services/gateway/config"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateShop(client pb.CorporationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input shop.CreateShopInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		input.CreatedBy = 1
		output, err := shop.CreateShop(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func UpdateShop(client pb.CorporationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input shop.UpdateShopInput
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
		output, err := shop.UpdateShop(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func DeleteShop(client pb.CorporationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input shop.DeleteShopInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = id
		output, err := shop.DeleteShop(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetShop(client pb.CorporationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input shop.GetShopInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = id
		output, err := shop.GetShop(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetListShop(client pb.CorporationClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input shop.GetListShopInput
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
		output, err := shop.GetListShop(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}
