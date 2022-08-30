package rest

import (
	"avatar/services/gateway/config"
	upload "avatar/services/gateway/infra/upload/respository"
	"avatar/services/gateway/pkg/util"
	pb "avatar/services/gateway/protos/center"
	"avatar/services/gateway/usecase/avatar"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateAvatar(
	client pb.CenterClient,
	upload upload.UploadClient,
	cfg *config.Environment,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input avatar.CreateAvatarInput
		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		startDate, err := time.Parse(time.RFC3339, input.StartDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		endDate, err := time.Parse(time.RFC3339, input.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if endDate.Sub(startDate) < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("start date must be less than end date")})
			return
		}
		if input.Vrm != nil {
			fileName, format := util.GenerateFileName(input.Vrm.Filename, input.Name)
			if format != ".VRM" {
				c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("the format of the image must be .vrm")})
				return
			}
			input.Vrm.Filename = fileName
		}
		if input.Image != nil {
			fileName, format := util.GenerateFileName(input.Image.Filename, input.Name)
			if format != ".PNG" && format != ".JPG" {
				c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("the format of the image must be PNG, JPG")})
				return
			}
			input.Image.Filename = fileName
		}
		output, err := avatar.CreateAvatar(&input, client, cfg, upload)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetAvatar(client pb.CenterClient, cfg *config.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input avatar.GetAvatarInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		input.Id = id
		output, err := avatar.GetAvatar(&input, client, cfg)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetListAvatar(client pb.CenterClient, cfg *config.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input avatar.GetListAvatarInput
		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if input.Gender > 2 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid gender, specified gender 1 is male and 2 is female"})
			return
		}
		output, err := avatar.GetListAvatar(&input, client, cfg)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func UpdateAvatar(
	client pb.CenterClient,
	upload upload.UploadClient,
	cfg *config.Environment,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input avatar.UpdateAvatarInput
		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		input.Id = id
		startDate, err := time.Parse(time.RFC3339, input.StartDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		endDate, err := time.Parse(time.RFC3339, input.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if endDate.Sub(startDate) < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("start date must be less than end date")})
			return
		}
		if input.Vrm != nil {
			fileName, format := util.GenerateFileName(input.Vrm.Filename, input.Name)
			if format != ".VRM" {
				c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("the format of the image must be .vrm")})
				return
			}
			input.Vrm.Filename = fileName
		}
		if input.Image != nil {
			fileName, format := util.GenerateFileName(input.Image.Filename, input.Name)
			if format != ".PNG" && format != ".JPG" {
				c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("the format of the image must be PNG, JPG")})
				return
			}
			input.Image.Filename = fileName
		}
		output, err := avatar.UpdateAvatar(&input, client, cfg, upload)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func DeleteAvatar(client pb.CenterClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input avatar.DeleteAvatarInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		input.Id = id
		output, err := avatar.DeleteAvatar(&input, client)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}
