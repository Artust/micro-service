package rest

import (
	"avatar/services/gateway/config"
	pb "avatar/services/gateway/protos/account_management"
	"avatar/services/gateway/usecase/account"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Message struct {
	Message string
}

func GetAccount(client pb.AccountManagementClient, cfg *config.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input account.GetAccountInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = id

		output, err := account.GetAccount(&input, client, cfg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetListAccount(client pb.AccountManagementClient, cfg *config.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input account.GetListAccountInput
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
		output, err := account.GetListAccount(&input, client, cfg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func CreateAccount(
	client pb.AccountManagementClient,
	cfg *config.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input account.CreateAccountInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if _, ok := pb.Gender_name[int32(input.Gender)]; !ok {
			c.JSON(http.StatusBadRequest, gin.H{"invalid gender": pb.Gender_name})
			return
		}
		output, err := account.CreateAccount(c, &input, client, cfg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func UpdateAccount(client pb.AccountManagementClient, cfg *config.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input account.UpdateAccountInput

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

		if input.Gender == 0 {
			if _, ok := pb.Gender_name[int32(input.Gender)]; !ok {
				c.JSON(http.StatusBadRequest, gin.H{"invalid gender": pb.Gender_name})
				return
			}
		}
		output, err := account.UpdateAccount(&input, client, cfg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func ChangePassword(client pb.AccountManagementClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input account.ChangePasswordInput
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

		err = account.ChangePassword(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, &Message{Message: config.ChangePassword})
	}
}

func DeactiveAccount(client pb.AccountManagementClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input account.DeactiveAccountInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = id
		err = account.DeactiveAccount(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, &Message{Message: config.DeactiveAccount})
	}
}

func ActiveAccount(client pb.AccountManagementClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input account.ActiveAccountInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = id

		err = account.ActiveAccount(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, &Message{Message: config.ActiveAccount})
	}
}
