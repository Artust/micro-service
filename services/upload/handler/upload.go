package handler

import (
	"avatar/services/upload/config"
	"avatar/services/upload/usecase/upload"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gin-gonic/gin"
)

func UploadToS3(s3 *session.Session, config *config.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input upload.UploadInput
		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		bucket := c.Param("bucket")
		input.Bucket = bucket
		err := upload.UploadToS3(&input, s3)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"url": fmt.Sprintf("/%s/%s", bucket, input.File.Filename),
		})
	}
}
