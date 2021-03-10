package models

import (
	"fmt"
	"gin-demo/app/config"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// PostID struct
type PostID struct {
	ID string `json:"id"`
}

// Post struct
type Post struct {
	PostID
	PostInfo
	TimeField
}

// PostInfo struct
type PostInfo struct {
	Slug           string  `json:"slug"`
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	Thumb          string  `json:"thumb"`
	Type           string  `json:"type"`
	Status         string  `json:"status"`
	Visibility     string  `json:"visibility"`
	AuthorID       int     `json:"author_id"`
	Author         *User   `gorm:"foreignKey:author_id;" json:"author"`
	Tags           []*Tag  `gorm:"many2many:coolpano_post_tags;joinForeignKey:post_id;" json:"tags"`
	Likes          []*User `gorm:"many2many:coolpano_post_likes;joinForeignKey:post_id;" json:"-"`
	Favorites      []*User `gorm:"many2many:coolpano_post_favorite;joinForeignKey:post_id;" json:"-"`
	LikesCount     int     `json:"likes_count" gorm:"-"`
	FavoritesCount int     `json:"favorites_count" gorm:"-"`
	Liked          bool    `json:"liked" gorm:"-"`
	Favorited      bool    `json:"favorited" gorm:"-"`
}

//TimeField struct
type TimeField struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostCommonView struct {
	PostID
	PostInfo
	Videos          []*Videos `gorm:"foreignKey:post_id;" json:"videos,omitempty"`
	Photos          []*Photos `gorm:"foreignKey:post_id;" json:"photos,omitempty"`
	Panorama        *Panorama `gorm:"foreignKey:post_id;" json:"panorama,omitempty"`
	AsteroidEntry   bool      `json:"asteroid_entry,omitempty"`
	Autorotate      bool      `json:"autorotate,omitempty"`
	AutorotateSpeed string    `json:"autorotate_speed,omitempty"`
	TimeField
}

//PostQ struct
type PostQ struct {
	Post
	PaginationQ
}

//PaginationQ gin handler query binding struct
type PaginationQ struct {
	Size  int         `form:"size" json:"size"`
	Page  int         `form:"page" json:"page"`
	Data  interface{} `json:"data" comment:"muster be a pointer of slice gorm.Model"` // save pagination list
	Total int64       `json:"total"`
}

//TableName Post
func (Post) TableName() string {
	return "coolpano_post"
}

//GetTotal 获取总数
func GetTotal(p *PaginationQ, queryTx *gorm.DB, list interface{}) (int64, error) {
	if p.Size < 1 {
		p.Size = 10
	}
	if p.Page < 1 {
		p.Page = 1
	}

	var total int64
	err := queryTx.Count(&total).Error
	if err != nil {
		return 0, err
	}
	offset := p.Size * (p.Page - 1)
	err = queryTx.Limit(p.Size).Offset(offset).Find(list).Error
	if err != nil {
		return 0, err
	}
	return total, err
}

//Search 查找
func (p PostQ) Search() (list *[]Post, total int64, err error) {
	list = &[]Post{}
	// var count int64
	tx := config.DB.Preload(clause.Associations).Find(&list)
	// tx := config.DB.Preload("Author").Preload("Tags").Find(&list)
	// tx := config.DB.Preload("Likes", func(db *gorm.DB) *gorm.DB {
	// 	return db.Table("coolpano_post_likes").Count(&count)
	// }).Find(&list)
	for _, p := range *list {
		p.LikesCount = 2
		p.FavoritesCount = 2
		fmt.Println(p.Slug)
		// fmt.Println(&list[k])

	}
	// fmt.Println(list)
	// count = config.DB.Model(&list).Association("Likes").Count()
	// fmt.Println(count)
	total, err = GetTotal(&p.PaginationQ, tx, list)
	return
}

//TypeChoices enum
func TypeChoices(t string) string {
	switch t {
	case "1":
		return "photo"
	case "2":
		return "video"
	case "3":
		return "photo360"
	default:
		return "photo"
	}
}

//StatusChoices enum
func StatusChoices(s string) string {
	switch s {
	case "1":
		return "accepted"
	case "2":
		return "processing"
	default:
		return "accepted"
	}
}

//VisibleChoices enum
func VisibleChoices(s string) string {
	switch s {
	case "1":
		return "public"
	case "2":
		return "private"
	default:
		return "public"
	}
}

//GetAllPost Fetch all post data
func (p Post) GetAllPost() (post *[]Post, err error) {
	// if err = config.DB.Find(&post).Scan(&result).Error; err != nil {
	post = &[]Post{}
	if err = config.DB.Preload(clause.Associations).Find(&post).Error; err != nil {
		return
	}
	return post, err
}

//GetPostByID ... Fetch only one user by Id
func (p Post) GetPostByID(id string) (post *Post, err error) {
	post = &Post{}
	if err = config.DB.Preload(clause.Associations).Where("id = ?", id).First(&post).Error; err != nil {
		return
	}
	post.LikesCount = len(post.Likes)
	post.FavoritesCount = len(post.Favorites)
	return post, err
}

//GetLikedOrFavorited 判断用户是否点赞/收藏此文章
func (p Post) GetLikedOrFavorited(list []*User, userid interface{}) bool {
	if userid != nil {
		for _, user := range list {
			if user.ID == userid {
				return true
			}
		}
		return false
	}
	return false
}

//GetPostWithDataBySlug func
func (p Post) GetPostWithDataBySlug(slug string) (post *PostCommonView, err error) {
	post = &PostCommonView{}
	if err = config.DB.Table(p.TableName()).Preload(clause.Associations).Where("slug = ?", slug).First(&post).Error; err != nil {
		return
	}
	post.Type = TypeChoices(post.Type)
	post.Status = StatusChoices(post.Status)
	post.Visibility = VisibleChoices(post.Visibility)
	post.LikesCount = len(post.Likes)
	post.FavoritesCount = len(post.Favorites)
	return post, err
}
