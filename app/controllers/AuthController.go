package controllers

import (
	password "gin-demo/app/helper"
	response "gin-demo/app/helper"
	token "gin-demo/app/helper"
	"gin-demo/app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// LoginView C
func LoginView(c *gin.Context) {
	// 用户发送用户名和密码过来
	var login models.LoginM
	query := &models.User{}
	if err := c.ShouldBind(&login); err != nil {
		response.JSON(c, http.StatusBadRequest, false, "无效的参数")
		return
	}
	// 将用户传入的用户名和密码和数据库中的进行比对
	user, err := query.GetUserByName(login.Username)
	if err != nil {
		response.JSON(c, http.StatusBadRequest, false, "用户名错误或者不存在")
		return
	}
	// 校验密码
	verified, _ := password.Verify(login.Password, user.Password)
	if verified {
		// 保存Session
		session := sessions.Default(c)
		session.Set("username", user.Username)
		session.Set("userid", user.ID)
		session.Save()
		// 生成Token
		tokenString, _ := token.GenerateToken(user.Username, 3*24*time.Hour)
		response.JSON(c, http.StatusOK, true, gin.H{"token": tokenString})
	} else {
		response.JSON(c, http.StatusBadRequest, false, "密码错误")
	}

}

// LogoutView C
func LogoutView(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	response.JSON(c, http.StatusOK, true, gin.H{"message": "User logout successfully"})
}
