package models

import "time"

type Photos struct {
	Slug      string    `json:"slug"`
	PostId    uint      `json:"post_id"`
	Url       string    `json:"url"`
	Width     string    `json:"width"`
	Height    string    `json:"height"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Photos) TableName() string {
	return "coolpano_post_photos"
}
