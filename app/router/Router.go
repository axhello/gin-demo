package router

import (
	"gin-demo/app/controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1/")
	{
		// UserController
		v1.GET("user", controllers.GetUsers)
		v1.POST("user", controllers.CreateUser)
		v1.GET("user/:id", controllers.GetUserByID)
		v1.PUT("user/:id", controllers.UpdateUser)
		v1.DELETE("user/:id", controllers.DeleteUser)
		// PostController
		v1.GET("post", controllers.GetPosts)
		v1.GET("post/:id", controllers.GetPostByID)
	}
	return r
}
