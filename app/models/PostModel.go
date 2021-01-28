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

type PostInfo struct {
	Slug           string `json:"slug"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Thumb          string `json:"thumb"`
	Type           string `json:"type"`
	Status         string `json:"status"`
	Visibility     string `json:"visibility"`
	LikesCount     int    `json:"likes_count" gorm:"-"`
	FavoritesCount int    `json:"favorites_count" gorm:"-"`
	Liked          bool   `json:"liked" gorm:"-"`
	Favorited      bool   `json:"favorited" gorm:"-"`
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
	AuthorId  int     `json:"author_id"`
	Author    *User   `gorm:"foreignKey:author_id;" json:"author"`
	Tags      []*Tag  `gorm:"many2many:coolpano_post_tags;" json:"tags"`
	Likes     []*User `gorm:"many2many:coolpano_post_likes;" json:"Likes"`
	Favorites []*User `gorm:"many2many:coolpano_post_favorite;" json:"Favorites"`
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

// # 判断用户是否点赞此文章
// def get_favorited(self, obj):
// 		user = self.get_request_user()
// 		return True if user in obj.favorite.all() else False

// 判断用户是否点赞此文章
func (p Post) GetLiked(userid interface{}, post *Post) bool {
	if userid != nil {
		for _, user := range post.Likes {
			if user.Id == userid {
				return true
			}
		}
		return false
	}
	return false
}

// 判断用户是否收藏此文章
func (p Post) GetFavorited(userid interface{}, post *Post) bool {
	if userid != nil {
		for _, user := range post.Favorites {
			if user.Id == userid {
				return true
			}
		}
		return false
	}
	return false
}

//GetPostPhotoBySlug
func (p Post) GetPostPhotoBySlug(slug string) (post *Post, err error) {
	post = &Post{}
	if err = config.DB.Preload(clause.Associations).Where("slug = ?", slug).First(&post).Error; err != nil {
		return
	}
	post.LikesCount = len(post.Likes)
	post.FavoritesCount = len(post.Favorites)
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
