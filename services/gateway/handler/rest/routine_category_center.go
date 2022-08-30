package rest

import (
	pb "avatar/services/gateway/protos/center"
	"avatar/services/gateway/usecase/routine_category_center"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRoutineCategory(client pb.CenterClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input routine_category_center.CreateRoutineCategoryInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := routine_category_center.CreateRoutineCategory(&input, client)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetRoutineCategory(client pb.CenterClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input routine_category_center.GetRoutineCategoryInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		input.Id = id
		output, err := routine_category_center.GetRoutineCategory(&input, client)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func GetListRoutineCategory(client pb.CenterClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input routine_category_center.GetListRoutineCategoryInput
		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		output, err := routine_category_center.GetListRoutineCategory(&input, client)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func UpdateRoutineCategory(client pb.CenterClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input routine_category_center.UpdateRoutineCategoryInput
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
		output, err := routine_category_center.UpdateRoutineCategory(&input, client)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}

func DeleteRoutineCategory(client pb.CenterClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input routine_category_center.DeleteRoutineCategoryInput
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		input.Id = id
		output, err := routine_category_center.DeleteRoutineCategory(&input, client)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, output)
	}
}
