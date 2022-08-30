package rest

import (
	"avatar/services/gateway/config"
	pb "avatar/services/gateway/protos/account_management"
	"fmt"
	"strconv"

	"avatar/services/gateway/usecase/permission"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPermission(client pb.AccountManagementClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input permission.GetPermissionInput
		id := c.Param("id")
		newId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = newId

		output, err := permission.GetPermission(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetListPermission(client pb.AccountManagementClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		output, err := permission.GetListPermission(client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func CreatePermission(client pb.AccountManagementClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input permission.CreatePermissionInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := permission.CreatePermission(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func UpdatePermission(client pb.AccountManagementClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input permission.UpdatePermissionInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		id := c.Param("id")
		newId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = newId

		output, err := permission.UpdatePermission(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func DeletePermission(client pb.AccountManagementClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input permission.DeletePermissionInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = id
		fmt.Println("Id: ", input.Id)

		err = permission.DeletePermission(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, &Message{Message: config.DeleteSuccess})
	}
}
