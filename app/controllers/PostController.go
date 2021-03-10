package controllers

import (
	response "gin-demo/app/helper"
	"net/http"

	"gin-demo/app/models"

	"github.com/gin-gonic/contrib/sessions"
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

//GetXML ... Get the pano xml
func GetXML(c *gin.Context) {
	slug := c.Params.ByName("slug")
	query := &models.Post{}
	post, err := query.GetPostWithDataBySlug(slug)
	if err != nil {
		response.JSON(c, http.StatusNotFound, false, err.Error())
	} else {
		siteMediaURL := "http://localhost:8000/media"
		c.Header("Content-Type", "application/xml")
		c.HTML(http.StatusOK, "normal.xml", gin.H{
			"site_media_url": siteMediaURL,
			"post":           post,
		})
	}
}

//GetPostBySlug ... Get the user by slug
func GetPostBySlug(c *gin.Context) {
	slug := c.Params.ByName("slug")
	session := sessions.Default(c)
	userid := session.Get("userid")
	query := &models.Post{}
	post, err := query.GetPostWithDataBySlug(slug)
	post.Liked = query.GetLikedOrFavorited(post.Likes, userid)
	post.Favorited = query.GetLikedOrFavorited(post.Favorites, userid)
	if err != nil {
		response.JSON(c, http.StatusNotFound, false, err.Error())
	} else {
		response.JSON(c, http.StatusOK, true, post)
	}
}
