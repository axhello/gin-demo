package models

import "time"

type Panorama struct {
	Slug      string    `json:"slug"`
	PostId    uint      `json:"post_id"`
	Original  string    `json:"original"`
	Thumb     string    `json:"thumb"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Panorama) TableName() string {
	return "coolpano_post_panorama"
}
