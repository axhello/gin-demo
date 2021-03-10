package router

import (
	"gin-demo/app/controllers"
	"gin-demo/app/middleware"
	"os"
	"path/filepath"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.LoadHTMLFiles(filepath.Join(os.Getenv("GOPATH"), "src/gin-demo/app/templates/coolpano/normal.xml"))
	v1 := r.Group("/api/v1")
	{
		// AuthController
		v1.POST("/auth/login", controllers.LoginView)
		v1.POST("/auth/signup", controllers.GetUsers)
		v1.POST("/auth/logout", controllers.LogoutView)

		// UserController
		v1.GET("/users", controllers.GetUsers)
		v1.POST("/user", controllers.CreateUser)
		v1.GET("/user", middleware.JWTAuth(), controllers.UserInfoView)
		v1.GET("/user/:id", controllers.GetUserByID)
		v1.PUT("/user/:id", controllers.UpdateUser)
		v1.DELETE("/user/:id", controllers.DeleteUser)

		// PostController
		v1.GET("/xml/:slug", controllers.GetXML)
		v1.GET("/posts", controllers.GetPosts)
		v1.GET("/post/:slug", controllers.GetPostBySlug)
	}
	return r
}
