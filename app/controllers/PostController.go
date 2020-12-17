package controllers

import (
	response "gin-demo/app/helper"
	"net/http"

	"gin-demo/app/models"

	"github.com/gin-gonic/gin"
)

//GetPosts ... Get all Posts
func GetPosts(c *gin.Context) {
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
	response.PaginationJSON(c, http.StatusOK, true, list, total, query.Page, query.Size)
}

func GetXml(c *gin.Context) {
	id := c.Params.ByName("id")
	query := &models.Post{}
	post, err := query.GetPostWithPanoramaById(id)
	if err != nil {
		response.JSON(c, http.StatusNotFound, false, err.Error())
	} else {
		site_media_url := "http://localhost:8000/media"
		c.Header("Content-Type", "application/xml")
		c.HTML(http.StatusOK, "normal.xml", gin.H{
			"site_media_url": site_media_url,
			"post":           post,
		})
	}

}

//GetPostByID ... Get the user by id
func GetPostByID(c *gin.Context) {
	id := c.Params.ByName("id")
	query := &models.Post{}
	post, err := query.GetPostById(id)
	if err != nil {
		response.JSON(c, http.StatusNotFound, false, err.Error())
	} else {
		response.JSON(c, http.StatusOK, true, post)
	}
}

//PhotosView
func PhotosView(c *gin.Context) {
	slug := c.Params.ByName("slug")
	query := &models.Post{}
	post, err := query.GetPostPhotoBySlug(slug)
	if err != nil {
		response.JSON(c, http.StatusNotFound, false, err.Error())
	} else {
		response.JSON(c, http.StatusOK, true, post)
	}
}
