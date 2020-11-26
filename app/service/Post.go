package service

import (
	"gin-demo/app/config"
	"gin-demo/app/models"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//GetAllPost Fetch all post data
func GetAllPost(post *[]models.Post) (err error) {
	// if err = config.DB.Find(&post).Scan(&result).Error; err != nil {
	if err = config.DB.Set("gorm:auto_preload", true).Find(&post).Error; err != nil {
		// if err = config.DB.Preload("Author").Preload("Tags").Preload("Likes").Preload("Favorites").Find(&post).Error; err != nil {
		return err
	}
	return nil
}

//GetPostByID ... Fetch only one user by Id
func GetPostByID(post *models.Post, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(&post).Error; err != nil {
		return err
	}
	return nil
}
