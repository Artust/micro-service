package middleware

import (
	"avatar/pkg/jwt"
	"avatar/services/gateway/config"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

func ignoreMethod(URI string, method string) bool {
	if URI == "/api/role/" && method == "GET" {
		return true
	}
	if URI == "/api/permission/" && method == "GET" {
		return true
	}
	return strings.HasPrefix(URI, "/api/auth")
}

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("URL: ", c.FullPath())
		if ignoreMethod(c.FullPath(), c.Request.Method) {
			return
		}
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := Response{
				Status:  false,
				Message: "Failed to process request",
				Errors:  "Not found token",
				Data:    nil,
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		cfg, err := config.Load()
		if err != nil {
			response := Response{
				Status:  false,
				Message: "Missing configuration",
				Errors:  "Not found config",
				Data:    nil,
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, response)
			return
		}
		idTokenHeader := strings.Split(authHeader, "Bearer ")
		if len(idTokenHeader) < 2 {
			response := Response{
				Status:  false,
				Message: "Must provide Authorization header with format `Bearer {token}",
				Errors:  "Invalid token",
				Data:    nil,
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		claims, err := jwt.VerifyToken(idTokenHeader[1], cfg.JwtSecretKey)
		if err != nil {
			log.Println(err)
			response := Response{
				Status:  false,
				Message: "Token is not valid",
				Errors:  err.Error(),
				Data:    nil,
			}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
		c.Set("claims", claims)
		log.Println("claims token: ", claims)
	}
}
