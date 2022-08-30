package rest

import (
	"avatar/services/gateway/config"
	pb "avatar/services/gateway/protos/center"
	"avatar/services/gateway/usecase/service_template_category"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateServiceTemplateCategory(centerClient pb.CenterClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input service_template_category.CreateCategoryInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := service_template_category.CreateServiceTemplateCategory(&input, centerClient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetServiceTemplateCategory(centerClient pb.CenterClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input service_template_category.GetCategoryInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		input.Id = id
		output, err := service_template_category.GetServiceTemplateCategory(&input, centerClient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetListServiceTemplateCategory(centerClient pb.CenterClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input service_template_category.GetListCategoryInput
		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		if input.Page == 0 {
			input.Page = config.Page
		}
		if input.PerPage == 0 {
			input.PerPage = config.PerPage
		}
		output, err := service_template_category.GetListServiceTemplateCategory(&input, centerClient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func UpdateServiceTemplateCategory(centerClient pb.CenterClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input service_template_category.UpdateCategoryInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		input.Id = id
		output, err := service_template_category.UpdateServiceTemplateCategory(&input, centerClient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func DeleteServiceTemplateCategory(centerClient pb.CenterClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input service_template_category.DeleteCategoryInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		input.Id = id
		output, err := service_template_category.DeleteServiceTemplateCategory(&input, centerClient)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}
