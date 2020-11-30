package controllers

import (
	"net/http"

	"gin-demo/app/models"

	"github.com/gin-gonic/gin"
)

func GetXml(c *gin.Context) {
	id := c.Params.ByName("id")
	query := &models.Post{}
	post, err := query.GetPostWithPanoramaById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"success": false,
			"message": err.Error(),
		})
	} else {
		// c.JSON(http.StatusOK, post)
		site_media_url := "http://localhost:8000/media"
		c.Header("Content-Type", "application/xml")
		c.HTML(http.StatusOK, "normal.xml", gin.H{
			"site_media_url": site_media_url,
			"post":           post,
		})
	}

}

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
		"success": true,
		"data":    list,
		"total":   total,
		"page":    query.Page,
		"size":    query.Size,
	})
}

//GetPostByID ... Get the user by id
func GetPostByID(c *gin.Context) {
	id := c.Params.ByName("id")
	query := &models.Post{}
	post, err := query.GetPostById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"success": false,
			"message": err.Error(),
		})
	} else {
		// fmt.Println(append(post, {"test": "user"}))
		c.JSON(http.StatusOK, post)
	}
}

//PhotosView
func PhotosView(c *gin.Context) {
	slug := c.Params.ByName("slug")
	query := &models.Post{}
	post, err := query.GetPostPhotoBySlug(slug)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"success": false,
			"message": err.Error(),
		})
	} else {
		// fmt.Println(append(post, {"test": "user"}))
		c.JSON(http.StatusOK, post)
	}
}
