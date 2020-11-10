//main.go
package main

import (
	"fmt"
	"gin-demo/app/config"
	"gin-demo/app/models"
	"gin-demo/app/router"
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	// gin.SetMode(gin.ReleaseMode)
	config.DB, err = gorm.Open("postgres", config.DbURL(config.BuildDBConfig()))
	config.DB.LogMode(true) // 开启sql日志
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer config.DB.Close()
	config.DB.AutoMigrate(&models.User{}, &models.Post{})
	r := router.SetupRouter()
	//running
	r.Run(":8090")

}

func init() {
	rand.Seed(time.Now().UnixNano())
}
