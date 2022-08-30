package rest

import (
	"avatar/services/gateway/config"
	upload "avatar/services/gateway/infra/upload/respository"
	"avatar/services/gateway/pkg/util"
	pb "avatar/services/gateway/protos/pos"
	"avatar/services/gateway/usecase/routine_pos"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetListRoutineByCategory(client pb.POSClient, cfg *config.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input routine_pos.GetListRoutineByCategoryInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		input.PosId = id
		output, err := routine_pos.GetListRoutineByCategory(&input, client, cfg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output.Results)
	}
}

func CreateRoutinePos(
	client pb.POSClient,
	upload upload.UploadClient,
	cfg *config.Environment,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input routine_pos.CreateRoutineInput
		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if input.PosId > 0 && input.ServiceTemplateId > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "routine has only 1 of 2 fields PosId or ServiceTemplateId"})
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
			c.JSON(http.StatusBadRequest, gin.H{"error": "start date must be less than end date"})
			return
		}
		output, err := routine_pos.CreateRoutine(&input, client, upload, cfg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func DeleteRoutinePos(client pb.POSClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input routine_pos.DeleteRoutineInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = id
		output, err := routine_pos.DeleteRoutine(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetListRoutinePos(client pb.POSClient, cfg *config.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input routine_pos.GetListRoutineInput
		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := routine_pos.GetListRoutine(&input, client, cfg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetRoutinePos(client pb.POSClient, cfg *config.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input routine_pos.GetRoutineInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = id
		output, err := routine_pos.GetRoutine(&input, client, cfg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func UpdateRoutinePos(client pb.POSClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input routine_pos.UpdateRoutineInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = id

		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if input.StartDate != "" && input.EndDate != "" {
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
				c.JSON(http.StatusBadRequest, gin.H{"error": "start date must be less than end date"})
				return
			}
		}
		output, err := routine_pos.UpdateRoutine(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}
