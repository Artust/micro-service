package util

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseIDFromURL(c *gin.Context) (int64, error) {
	id := c.Param("id")
	newId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"input id must be numeric, error": err.Error()})
		return 0, err
	}
	return newId, nil
}
