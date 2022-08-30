package logger

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				/*
					var (
						stack  = make([]byte, defaultStackSize)
						length = runtime.Stack(stack, true)
					)
				*/
				log.Error().Stack().Err(err).Msg("")

				c.AbortWithStatusJSON(http.StatusInternalServerError, &gin.H{
					"errors": []string{"unexpected error"},
				})
			}
		}()
		c.Next()
	}
}
