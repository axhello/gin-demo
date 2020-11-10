package controllers

import (
	"net/http"

	"gin-demo/app/models"
	"gin-demo/app/service"

	"github.com/gin-gonic/gin"
)

//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	// var user []models.User
	// err := service.GetAllUsers(&user)
	user := &models.User{}
	list, err := user.GetAllUsers()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"status":  "error",
			"message": "Not Found!",
		})
	} else {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"status": "success",
			"data":   list,
		})
	}
}

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": "非法数据格式",
		})
		return
	}
	list, err := user.CreateUser()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "error",
			"message": err.Error(),
		})
	} else {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"status": "success",
			"data":   list,
		})
	}
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	user := &models.User{}
	list, err := user.GetUserByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"status":  "error",
			"message": "Not Found!",
		})
	} else {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"code":   http.StatusOK,
			"status": "success",
			"data":   list,
		})
	}
}

//UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	err := service.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = service.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Params.ByName("id")
	err := service.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
