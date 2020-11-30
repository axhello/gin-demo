package router

import (
	"gin-demo/app/controllers"
	"gin-demo/app/middleware"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLFiles(filepath.Join(os.Getenv("GOPATH"), "src/gin-demo/app/templates/coolpano/normal.xml"))
	v1 := r.Group("/api/v1")
	{
		// AuthController
		v1.POST("/auth/login", controllers.LoginView)
		v1.POST("/auth/signup", controllers.GetUsers)
		v1.POST("/auth/logout", controllers.GetUsers)

		// UserController
		v1.GET("/users", controllers.GetUsers)
		v1.POST("/user", controllers.CreateUser)
		v1.GET("/user/:id", controllers.GetUserByID)
		v1.PUT("/user/:id", controllers.UpdateUser)
		v1.DELETE("/user/:id", controllers.DeleteUser)
		// PostController
		v1.GET("/xml/:id", middleware.JWTAuthMiddleware(), controllers.GetXml)
		v1.GET("/posts", controllers.GetPosts)
		v1.GET("/post/:id", controllers.GetPostByID)
		v1.GET("/photos/:slug", controllers.PhotosView)
		v1.GET("/videos/:slug", controllers.GetPostByID)
		v1.GET("/photo-360s/:slug", controllers.GetPostByID)
	}
	return r
}
