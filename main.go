//main.go
package main

import (
	"fmt"
	"gin-demo/app/config"
	"gin-demo/app/router"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var err error

func main() {
	// gin.SetMode(gin.ReleaseMode)
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
