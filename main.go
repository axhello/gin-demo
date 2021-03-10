//main.go
package main

import (
	"fmt"
	"gin-demo/app/config"
	"gin-demo/app/router"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var err error

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	gin.SetMode(os.Getenv("MODE"))
	config.DB, err = gorm.Open(postgres.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 开启sql日志
	})
	if err != nil {
		fmt.Println("Status:", err)
	}
	// config.DB.AutoMigrate(&models.User{}, &models.Post{})
	r := router.SetupRouter()
	r.Run(":8090")
}
