package helper

import (
	"github.com/gin-gonic/gin"
)

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
