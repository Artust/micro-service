package rest

import (
	"avatar/services/gateway/config"
	upload "avatar/services/gateway/infra/upload/respository"
	"avatar/services/gateway/usecase/upload_file"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadFile(
	upload upload.UploadClient,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		bucket := c.Param("bucket")
		var input upload_file.UploadFileInput
		listBucket := config.Bucket
		for key, val := range listBucket {
			if key == len(listBucket)-1 && val != bucket {
				c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("no buckets are valid").Error()})
				return
			}
			if val == bucket {
				input.Bucket = bucket
				break
			}
		}
		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := upload_file.UploadFile(&input, upload)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}
