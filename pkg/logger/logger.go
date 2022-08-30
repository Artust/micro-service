package logger

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func SetLogger() gin.HandlerFunc {
	sublog := log.Logger

	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		statusCode := c.Writer.Status()
		dataLength := c.Writer.Size()
		if dataLength < 0 {
			dataLength = 0
		}

		dumplogger := sublog.With().
			Str("remoteaddr", c.ClientIP()).
			Str("realip", c.Request.Header.Get("X-Real-IP")).
			Str("forwardedfor", c.Request.Header.Get("X-Forwarded-For")).
			Str("method", c.Request.Method).
			Int64("reqlength", c.Request.ContentLength).
			Str("path", path).
			Str("query", query).
			Int("status", statusCode).
			Int("size", dataLength).
			Str("ref", c.Request.Referer()).
			Str("ua", c.Request.UserAgent()).
			Str("cookie", c.Request.Header.Get("Cookie")).
			Dur("reqtime", latency).
			Str("host", c.Request.Host).
			Str("reqid", c.Request.Header.Get("X-Request-ID")).
			Logger()

		if len(c.Errors) > 0 {
			dumplogger.Error().Msg(c.Errors.ByType(gin.ErrorTypePrivate).String())
		} else {
			if statusCode >= http.StatusInternalServerError {
				dumplogger.Error().Send()
			} else if statusCode >= http.StatusBadRequest {
				dumplogger.Warn().Send()
			} else {
				dumplogger.Info().Send()
			}
		}
	}
}
