package models

import "time"

//Videos struct
type Videos struct {
	Slug      string    `json:"slug"`
	PostID    uint      `json:"post_id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	Cover     string    `json:"cover"`
	Width     string    `json:"width"`
	Height    string    `json:"height"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//TableName Videos
func (Videos) TableName() string {
	return "coolpano_post_videos"
}
