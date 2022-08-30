package rest

import (
	"avatar/services/gateway/config"
	upload "avatar/services/gateway/infra/upload/respository"
	pb "avatar/services/gateway/protos/center"
	"avatar/services/gateway/usecase/service_template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateServiceTemplate(
	centerClient pb.CenterClient,
	upload upload.UploadClient,
	cfg *config.Environment,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input service_template.CreateServiceTemplateInput
		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := service_template.CreateServiceTemplate(&input, centerClient, upload, cfg)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func UpdateServiceTemplate(
	centerClient pb.CenterClient,
	cfg *config.Environment,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input service_template.UpdateServiceTemplateInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = id
		output, err := service_template.UpdateServiceTemplate(&input, centerClient, cfg)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func DeleteServiceTemplate(centerClient pb.CenterClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input service_template.DeleteServiceTemplateInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.Id = id
		output, err := service_template.DeleteServiceTemplate(&input, centerClient)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetServiceTemplate(centerClient pb.CenterClient, cfg *config.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input service_template.GetServiceTemplateInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
			return
		}
		input.ID = id
		output, err := service_template.GetServiceTemplate(&input, centerClient, cfg)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetListServiceTemplate(centerClient pb.CenterClient, cfg *config.Environment) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input service_template.GetListServiceTemplateInput
		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := service_template.GetListServiceTemplate(&input, centerClient, cfg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}
