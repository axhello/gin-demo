package models

import "time"

type Videos struct {
	Slug      string    `json:"slug"`
	PostId    uint      `json:"post_id"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	Cover     string    `json:"cover"`
	Width     string    `json:"width"`
	Height    string    `json:"height"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Videos) TableName() string {
	return "coolpano_post_videos"
}
