package rest

import (
	"avatar/services/gateway/config"
	pb "avatar/services/gateway/protos/account_management"
	"strconv"

	"avatar/services/gateway/usecase/account_role"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAccountRole(client pb.AccountManagementClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input account_role.GetRoleInput
		id := c.Param("id")
		newId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = newId
		output, err := account_role.GetRole(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetListAccountRole(client pb.AccountManagementClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		output, err := account_role.GetListAccountRole(client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func CreateAccountRole(client pb.AccountManagementClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input account_role.CreateRoleInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := account_role.CreateRole(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func UpdateAccountRole(client pb.AccountManagementClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input account_role.UpdateRoleInput
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

		output, err := account_role.UpdateRole(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func DeleteAccountRole(client pb.AccountManagementClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input account_role.DeleteRoleInput
		id := c.Param("id")
		newId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = newId
		err = account_role.DeleteRole(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, &Message{Message: config.DeleteSuccess})
	}
}
