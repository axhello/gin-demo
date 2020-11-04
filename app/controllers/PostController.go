package controllers

import (
	"net/http"

	"gin-demo/app/models"
	"gin-demo/app/service"

	"github.com/gin-gonic/gin"
)

//GetPosts ... Get all users
func GetPosts(c *gin.Context) {
	// var post []models.Post
	// err := service.GetAllPost(&post)
	// if err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"code":    http.StatusNotFound,
	// 		"status":  "error",
	// 		"message": "data not found",
	// 	})
	// } else {
	// 	c.JSON(http.StatusOK, post)
	// }
	query := &models.PostQ{}
	err := c.ShouldBindQuery(query) //开始绑定url-query 参数到结构体
	if err != nil {
		return
	}
	list, total, err := query.Search() //开始mysql 业务搜索查询
	if err != nil {
		return
	}
	//返回数据开始拼装分页json
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   list,
		"total":  total,
		"page":   query.Page,
		"size":   query.Size,
	})
}

//GetPostByID ... Get the user by id
func GetPostByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var post models.Post
	err := service.GetPostByID(&post, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"status":  "error",
			"message": "data not found",
		})
	} else {
		// fmt.Println(append(post, {"test": "user"}))
		c.JSON(http.StatusOK, post)
	}
}