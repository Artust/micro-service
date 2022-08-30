package rest

import (
	"avatar/services/gateway/config"
	pb "avatar/services/gateway/protos/pos"
	"avatar/services/gateway/usecase/routine_category_pos"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRoutineCategoryPOS(client pb.POSClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input routine_category_pos.CreateRoutineCategoryInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := routine_category_pos.CreateRoutineCategory(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetRoutineCategoryPOS(client pb.POSClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input routine_category_pos.GetRoutineCategoryInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		input.Id = id
		output, err := routine_category_pos.GetRoutineCategory(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetListRoutineCategoryPOS(client pb.POSClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input routine_category_pos.GetListRoutineCategoryInput
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
		output, err := routine_category_pos.GetListRoutineCategory(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func UpdateRoutineCategoryPOS(client pb.POSClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input routine_category_pos.UpdateRoutineCategoryInput
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
		output, err := routine_category_pos.UpdateRoutineCategory(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func DeleteRoutineCategoryPOS(client pb.POSClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input routine_category_pos.DeleteRoutineCategoryInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		input.Id = id
		output, err := routine_category_pos.DeleteRoutineCategory(&input, client)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}
