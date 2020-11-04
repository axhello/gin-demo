package models

type Tag struct {
	Id   string `json:"id"`
	Slug string `json:"slug"`
	Text string `json:"text"`
}

func (Tag) TableName() string {
	return "coolpano_tag"
}
