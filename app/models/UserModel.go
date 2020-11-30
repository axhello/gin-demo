package models

import (
	"errors"
	"fmt"
	password "gin-demo/app/helper"

	"gin-demo/app/config"
)

type UserId struct {
	Id uint `gorm:"primary_key;AUTO_INCREMENT"  json:"id"` // 自增id
}

type PassWord struct {
	Password string `json:"-"`
}

type UserInfo struct {
	Slug        string `json:"slug"`
	IsSuperuser bool   `json:"is_superuser"`
	// Nickname string `json:"nickname"`
	// Email    string `json:"email"`    //邮箱
	// Avatar   string `json:"avatar"`   //头像
	Sex string `json:"sex"` //性别
	// Cover    string `json:"cover"`    //封面
	// Weibo    string `json:"weibo"`    //微博
	// Facebook string `json:"facebook"` //脸书
	// Twitter  string `json:"twitter"`  //推特
	// Website  string `json:"website"`  //网站
}

type LoginM struct {
	Username string `json:"username"`
	PassWord
}

type UserView struct {
	UserId
	LoginM
	UserInfo
}

type User struct {
	UserId
	LoginM
	UserInfo
}

func (User) TableName() string {
	return "coolpano_user"
}

//GetAllUsers Fetch all user data
func (u *User) GetAllUsers() (UserView []UserView, err error) {
	// user = &[]User{}
	if err = config.DB.Table(u.TableName()).Find(&UserView).Error; err != nil {
		return
	}
	return
}

//CreateUser ... Insert New data
func (u *User) CreateUser() (user *User, err error) {
	// if err = u.Encrypt(); err != nil {
	// 	return
	// }
	pwd, err := password.Encode(u.Password, "", 0)
	if err != nil {
		err = errors.New("密码加密失败！")
		return
	}
	u.Password = string(pwd)
	u.Slug = RandStr(10)
	u.Sex = "0"
	fmt.Println(u)
	// check 用户名
	var count int64
	config.DB.Table(u.TableName()).Where("username = ?", u.Username).Count(&count)
	if count > 0 {
		err = errors.New("账户已存在！")
		return
	}
	if err = config.DB.Table(u.TableName()).Create(u).Error; err != nil {
		return
	}
	return user, err
}

//GetUserByID ... Fetch only one user by Id
func (u *User) GetUserByID(id string) (user *User, err error) {
	user = &User{}
	if err = config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return
	}
	return user, err
}

//GetUserByName ... Fetch only one user by username
func (u *User) GetUserByName(username string) (user *User, err error) {
	user = &User{}
	if err = config.DB.Table(u.TableName()).Where("username = ?", username).First(&user).Error; err != nil {
		return
	}
	return user, err
}

//UpdateUser ... Update user
func (u *User) UpdateUser(id string) (user *User, err error) {
	// config.DB.Save(user)
	if err = config.DB.Save(user).Error; err != nil {
		return
	}
	return user, err
}

//DeleteUser ... Delete user
func (u *User) DeleteUser(user *User, id string, err error) {
	if err := config.DB.Where("id = ?", id).Delete(user).Error; err != nil {
		return
	}
}
