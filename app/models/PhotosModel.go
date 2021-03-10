package models

import "time"

// Photos struct
type Photos struct {
	Slug      string    `json:"slug"`
	PostID    uint      `json:"post_id"`
	URL       string    `json:"url"`
	Width     string    `json:"width"`
	Height    string    `json:"height"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName Photos
func (Photos) TableName() string {
	return "coolpano_post_photos"
}
