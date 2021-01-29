package models

import (
	"fmt"
	"gin-demo/app/config"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PostId struct {
	Id string `json:"id"`
}

type Post struct {
	PostId
	PostInfo
	TimeField
}

type PostInfo struct {
	Slug           string  `json:"slug"`
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	Thumb          string  `json:"thumb"`
	Type           string  `json:"type"`
	Status         string  `json:"status"`
	Visibility     string  `json:"visibility"`
	AuthorId       int     `json:"author_id"`
	Author         *User   `gorm:"foreignKey:author_id;" json:"author"`
	Tags           []*Tag  `gorm:"many2many:coolpano_post_tags;joinForeignKey:post_id;" json:"tags"`
	Likes          []*User `gorm:"many2many:coolpano_post_likes;joinForeignKey:post_id;" json:"-"`
	Favorites      []*User `gorm:"many2many:coolpano_post_favorite;joinForeignKey:post_id;" json:"-"`
	LikesCount     int     `json:"likes_count" gorm:"-"`
	FavoritesCount int     `json:"favorites_count" gorm:"-"`
	Liked          bool    `json:"liked" gorm:"-"`
	Favorited      bool    `json:"favorited" gorm:"-"`
}

type TimeField struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostPanoramaView struct {
	PostId
	PostInfo
	AsteroidEntry   bool      `json:"asteroid_entry"`
	Autorotate      bool      `json:"autorotate"`
	AutorotateSpeed string    `json:"autorotate_speed"`
	Panorama        *Panorama `gorm:"foreignKey:post_id;" json:"panorama"`
	TimeField
}

type PostPhotosView struct {
	PostId
	PostInfo
	Photos []*Photos `gorm:"foreignKey:post_id;" json:"photos"`
	TimeField
}

type PostVideosView struct {
	PostId
	PostInfo
	Videos []*Videos `gorm:"foreignKey:post_id;" json:"videos"`
	TimeField
}

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

func (Post) TableName() string {
	return "coolpano_post"
}

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
func (p Post) GetPostById(id string) (post *Post, err error) {
	post = &Post{}
	if err = config.DB.Preload(clause.Associations).Where("id = ?", id).First(&post).Error; err != nil {
		return
	}
	post.LikesCount = len(post.Likes)
	post.FavoritesCount = len(post.Favorites)
	return post, err
}

// 判断用户是否点赞/收藏此文章
func (p Post) GetLikedOrFavorited(userid interface{}, list []*User) bool {
	if userid != nil {
		for _, user := range list {
			if user.Id == userid {
				return true
			}
		}
		return false
	}
	return false
}
func TYPE_CHOICES(t string) string {
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
func STATUS_CHOICES(s string) string {
	switch s {
	case "1":
		return "accepted"
	case "2":
		return "processing"
	default:
		return "accepted"
	}
}
func VISIBLE_CHOICES(s string) string {
	switch s {
	case "1":
		return "public"
	case "2":
		return "private"
	default:
		return "public"
	}
}

//GetPostWithPhotoBySlug
func (p Post) GetPostWithPhotoBySlug(slug string) (post *PostPhotosView, err error) {
	post = &PostPhotosView{}
	if err = config.DB.Table(p.TableName()).Preload(clause.Associations).Where("slug = ?", slug).First(&post).Error; err != nil {
		return
	}
	post.Type = TYPE_CHOICES(post.Type)
	post.Status = STATUS_CHOICES(post.Status)
	post.Visibility = VISIBLE_CHOICES(post.Visibility)
	post.LikesCount = len(post.Likes)
	post.FavoritesCount = len(post.Favorites)
	return post, err
}

//GetPostWithVideoBySlug
func (p Post) GetPostWithVideoBySlug(slug string) (post *PostVideosView, err error) {
	post = &PostVideosView{}
	if err = config.DB.Table(p.TableName()).Preload(clause.Associations).Where("slug = ?", slug).First(&post).Error; err != nil {
		return
	}
	post.Type = TYPE_CHOICES(post.Type)
	post.Status = STATUS_CHOICES(post.Status)
	post.Visibility = VISIBLE_CHOICES(post.Visibility)
	post.LikesCount = len(post.Likes)
	post.FavoritesCount = len(post.Favorites)
	return post, err
}

//GetPostWithPanoramaBySlug
func (p Post) GetPostWithPanoramaBySlug(id string) (post *PostPanoramaView, err error) {
	post = &PostPanoramaView{}
	if err = config.DB.Table(p.TableName()).Preload(clause.Associations).Where("slug = ?", id).First(&post).Error; err != nil {
		return
	}
	post.Type = TYPE_CHOICES(post.Type)
	post.Status = STATUS_CHOICES(post.Status)
	post.Visibility = VISIBLE_CHOICES(post.Visibility)
	post.LikesCount = len(post.Likes)
	post.FavoritesCount = len(post.Favorites)
	return post, err
}
