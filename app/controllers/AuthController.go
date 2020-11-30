package controllers

import (
	"fmt"
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
			"status":  "error",
			"message": "无效的参数",
		})
		return
	}
	// 将用户传入的用户名和密码和数据库中的进行比对
	user, err := query.GetUserByName(login.Username)
	if err != nil {
		fmt.Println("get user from db by name error")
		return
	}
	// 校验用户名和密码是否正确
	if user.Username == "q1mi" {
		// 生成Token
		tokenString, _ := token.GenerateToken(user.Username, 3*24*time.Hour)
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"status": "success",
			"data":   gin.H{"token": tokenString},
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg":  "鉴权失败",
	})
}
