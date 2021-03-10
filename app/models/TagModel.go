package models

// Tag struct
type Tag struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
	Text string `json:"text"`
}

//TableName Tag
func (Tag) TableName() string {
	return "coolpano_tag"
}
