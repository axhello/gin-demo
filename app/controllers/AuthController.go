package controllers

import (
	password "gin-demo/app/helper"
	token "gin-demo/app/helper"
	"gin-demo/app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginView(c *gin.Context) {
	// 用户发送用户名和密码过来
	var login models.LoginM
	query := &models.User{}
	if err := c.ShouldBind(&login); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"success": false,
			"message": "无效的参数",
		})
		return
	}
	// 将用户传入的用户名和密码和数据库中的进行比对
	user, err := query.GetUserByName(login.Username)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"success": false,
			"message": "用户名错误或者不存在",
		})
		return
	}
	// 校验密码
	verified, _ := password.Verify(login.Password, user.Password)
	if verified {
		// 生成Token
		tokenString, _ := token.GenerateToken(user.Username, 3*24*time.Hour)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
			"data":    gin.H{"token": tokenString},
		})
		return
	} else {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"success": false,
			"message": "密码错误",
		})
		return
	}

}
