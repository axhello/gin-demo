package models

type User struct {
	Id       uint   `json:"id"`
	Slug     string `json:"slug"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

func (User) TableName() string {
	return "coolpano_user"
}

// func (CustomUser) TableName() string {
// 	return "coolpano_user"
// }
