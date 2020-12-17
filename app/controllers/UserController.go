package controllers

import (
	response "gin-demo/app/helper"
	"gin-demo/app/models"
	"gin-demo/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserInfoView(c *gin.Context) {
	user := &models.User{}
	username, _ := c.Get("username")
	list, err := user.GetUserByName(username)
	if err != nil {
		response.JSON(c, http.StatusNotFound, false, err.Error())
	} else {
		response.JSON(c, http.StatusOK, true, list)
	}
}

//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	// var user []models.User
	// err := service.GetAllUsers(&user)
	user := &models.User{}
	list, err := user.GetAllUsers()
	if err != nil {
		response.JSON(c, http.StatusNotFound, false, err.Error())
	} else {
		response.JSON(c, http.StatusOK, true, list)
	}
}

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		response.JSON(c, http.StatusBadRequest, false, "非法数据格式")
		return
	}
	list, err := user.CreateUser()
	if err != nil {
		response.JSON(c, http.StatusNotFound, false, err.Error())
	} else {
		response.JSON(c, http.StatusOK, true, list)
	}
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	user := &models.User{}
	list, err := user.GetUserByID(id)
	if err != nil {
		response.JSON(c, http.StatusNotFound, false, err.Error())
	} else {
		response.JSON(c, http.StatusOK, true, list)
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
