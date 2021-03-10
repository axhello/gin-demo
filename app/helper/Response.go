package helper

import (
	"github.com/gin-gonic/gin"
)

// JSON response
func JSON(c *gin.Context, httpStatus int, success bool, data interface{}) {
	if success {
		c.AbortWithStatusJSON(httpStatus, gin.H{
			"code":    httpStatus,
			"success": success,
			"data":    data,
		})
	} else {
		c.AbortWithStatusJSON(httpStatus, gin.H{
			"code":    httpStatus,
			"success": success,
			"message": data,
		})

	}

}

// PaginationJSON response
func PaginationJSON(c *gin.Context, httpStatus int, success bool, data interface{}, total int64, page int, size int) {
	c.AbortWithStatusJSON(httpStatus, gin.H{
		"code":    httpStatus,
		"success": success,
		"data":    data,
		"total":   total,
		"page":    page,
		"size":    size,
	})
}
