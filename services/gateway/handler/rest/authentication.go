package rest

import (
	"avatar/services/gateway/config"
	pb "avatar/services/gateway/protos/account_management"
	"avatar/services/gateway/usecase/account"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(client pb.AccountManagementClient, cfg *config.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input account.LoginInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := account.Login(&input, client, cfg)
		if err != nil {
			if _, ok := config.Auth_message_code[err.Error()]; ok {
				c.JSON(http.StatusUnauthorized, gin.H{
					"messageCode": config.Auth_message_code[err.Error()],
					"data":        nil,
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"messageCode": config.Auth_message_code["unspecified"],
				"data":        nil,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"messageCode": config.Auth_message_code["ok"],
			"data":        output,
		})
	}
}

func ForgotPassword(client pb.AccountManagementClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input account.ForgotPasswordInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		err := account.ForgotPassword(&input, client)
		if err != nil {
			if _, ok := config.Auth_message_code[err.Error()]; ok {
				c.JSON(http.StatusUnauthorized, gin.H{
					"messageCode": config.Auth_message_code[err.Error()],
					"data":        nil,
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"messageCode": config.Auth_message_code["unspecified"],
				"data":        nil,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"messageCode": config.Auth_message_code["ok"],
			"data":        nil,
		})
	}
}

func ResetPassword(client pb.AccountManagementClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input account.ResetPasswordInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		resetPwdToken := c.Query("token")
		fmt.Println("Reset Password Token: ", resetPwdToken)
		input.ResetPasswordToken = resetPwdToken
		err := account.ResetPassword(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, &Message{Message: config.ChangePassword})
	}
}
