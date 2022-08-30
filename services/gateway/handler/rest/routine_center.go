package rest

import (
	"avatar/services/gateway/config"
	upload "avatar/services/gateway/infra/upload/respository"
	"avatar/services/gateway/pkg/util"
	pb "avatar/services/gateway/protos/center"
	"avatar/services/gateway/usecase/routine_center"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetListRoutineCenterByCategory(client pb.CenterClient, cfg *config.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input routine_center.GetListRoutineByCategoryInput
		// serviceTemplateId, err := strconv.ParseInt(c.Param("id"), 10, 64)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
		// 	return
		// }
		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		// input.ServiceTemplateId = serviceTemplateId
		output, err := routine_center.GetListRoutineByCategory(&input, client, cfg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output.Results)
	}
}

func GetListRoutineCenter(centerClient pb.CenterClient, cfg *config.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input routine_center.GetListRoutineInput
		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := routine_center.GetListRoutineCenter(&input, centerClient, cfg)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetRoutineCenter(centerClient pb.CenterClient, cfg *config.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input routine_center.GetRoutineCenterInput
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		input.ID = id
		output, err := routine_center.GetRoutineCenter(&input, centerClient, cfg)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func CreateRoutineCenter(
	centerClient pb.CenterClient,
	upload upload.UploadClient,
	cfg *config.Environment,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input routine_center.CreateRoutineCenterInput
		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		_, format := util.GenerateFileName(input.AnimationFile.Filename, input.Name)
		if format != ".DAT" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "the format of the animation file must be dat"})
			return
		}
		_, format = util.GenerateFileName(input.SoundFile.Filename, input.Name)
		if format != ".MP3" && format != ".WAV" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "the format of the sound file must be MP3, WAV"})
			return
		}
		if input.ImageFile != nil {
			_, format = util.GenerateFileName(input.ImageFile.Filename, input.Name)
			if format != ".PNG" && format != ".JPG" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "the format of the image must be PNG, JPG"})
				return
			}
		}
		startDate, err := time.Parse(time.RFC3339, input.StartDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		endDate, err := time.Parse(time.RFC3339, input.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if endDate.Sub(startDate) < 0 {
			c.JSON(http.StatusBadRequest, errors.New("start date must be less than end date"))
			return

		}
		output, err := routine_center.CreateRoutineCenter(&input, centerClient, upload, cfg)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func DeleteRoutineCenter(centerClient pb.CenterClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input routine_center.DeleteRoutineCenterInput
		id, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		input.ID = id
		output, err := routine_center.DeleteRoutineCenter(&input, centerClient)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func UpdateRoutineCenter(client pb.CenterClient, cfg *config.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input routine_center.UpdateRoutineInput
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
		startDate, err := time.Parse(time.RFC3339, input.StartDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, errors.New("error when parse start date"))
			return
		}
		endDate, err := time.Parse(time.RFC3339, input.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, errors.New("error when parse end date"))
			return
		}
		if endDate.Sub(startDate) < 0 {
			c.JSON(http.StatusBadRequest, errors.New("start date must be less than end date"))
			return
		}
		output, err := routine_center.UpdateRoutine(&input, client, cfg)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}
