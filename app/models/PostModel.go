package models

import (
	"fmt"
	"gin-demo/app/config"
	"time"
)

type PostId struct {
	Id string `json:"id"`
}

type PostInfo struct {
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Thumb       string `json:"thumb"`
	Type        string `json:"type"`
	Status      string `json:"status"`
	Visibility  string `json:"visibility"`
	LikesCount  uint   `gorm:"-"`
}

type TimeField struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostTag struct {
	AuthorId int   `json:"author_id"`
	Author   *User `gorm:"foreignKey:author_id;" json:"author"`
	Tags     []Tag `gorm:"many2many:coolpano_post_tags;" json:"tags"`
	// Likes     []*User `gorm:"many2many:coolpano_post_likes;" json:"likes"`
	// Favorites []*User `gorm:"many2many:coolpano_post_favorite;" json:"favorites"`
}

type PostPanoramaView struct {
	PostId
	PostInfo
	AsteroidEntry   string    `json:"asteroid_entry"`
	Autorotate      string    `json:"autorotate"`
	AutorotateSpeed string    `json:"autorotate_speed"`
	Panorama        *Panorama `gorm:"foreignKey:post_id;" json:"panorama"`
	TimeField
}

type Post struct {
	PostId
	PostInfo
	AuthorId int     `json:"author_id"`
	Author   *User   `gorm:"foreignKey:author_id;" json:"author"`
	Tags     []*Tag  `gorm:"many2many:coolpano_post_tags;" json:"tags"`
	Likes    []*User `gorm:"many2many:coolpano_post_likes;" json:"likes"`
	TimeField
}

type PostQ struct {
	Post
	PaginationQ
	FromTime string `form:"from_time"` //搜索开始时间
	ToTime   string `form:"to_time"`   //搜索结束时候
}

func (Post) TableName() string {
	return "coolpano_post"
}

func (p PostQ) Search() (list *[]Post, total uint, err error) {
	list = &[]Post{}
	tx := config.DB.Set("gorm:auto_preload", true).Find(&list)
	if p.FromTime != "" && p.ToTime != "" {
		tx = tx.Where("`created_at` BETWEEN ? AND ?", p.FromTime, p.ToTime)
	}
	total, err = crudAll(&p.PaginationQ, tx, list)
	return
}

//GetAllPost Fetch all post data
func GetAllPost(post *[]Post) (err error) {
	// if err = config.DB.Find(&post).Scan(&result).Error; err != nil {
	if err = config.DB.Set("gorm:auto_preload", true).Find(&post).Error; err != nil {
		return err
	}
	return nil
}

//GetPostByID ... Fetch only one user by Id
func (p Post) GetPostById(id string) (post *Post, err error) {
	post = &Post{}
	var count int64
	config.DB.Table(p.TableName()).Preload("Likes").Where("id = ?", id).Count(&count)
	fmt.Println(count)
	if err = config.DB.Set("gorm:auto_preload", true).Where("id = ?", id).First(&post).Error; err != nil {
		return
	}
	return post, err
}

//GetPostByID ... Fetch only one user by Id
func (p Post) GetPostWithPanoramaById(id string) (post *PostPanoramaView, err error) {
	post = &PostPanoramaView{}
	if err = config.DB.Table(p.TableName()).Preload("Panorama").Where("id = ?", id).First(&post).Error; err != nil {
		return
	}
	return post, err
}
