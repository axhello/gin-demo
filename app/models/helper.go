package models

import (
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//PaginationQ gin handler query binding struct
type PaginationQ struct {
	Ok    bool        `json:"ok"`
	Size  uint        `form:"size" json:"size"`
	Page  uint        `form:"page" json:"page"`
	Data  interface{} `json:"data" comment:"muster be a pointer of slice gorm.Model"` // save pagination list
	Total uint        `json:"total"`
}

//SearchAll optimized pagination method for gorm
func (p *PaginationQ) SearchAll(queryTx *gorm.DB) (data *PaginationQ, err error) {
	//99999 magic number for get all list without pagination
	if p.Size == 9999 || p.Size == 99999 {
		err = queryTx.Find(p.Data).Error
		p.Ok = err == nil
		return p, err
	}

	if p.Size < 1 {
		p.Size = 10
	}
	if p.Page < 1 {
		p.Page = 1
	}
	offset := p.Size * (p.Page - 1)
	err = queryTx.Count(&p.Total).Error
	if err != nil {
		return p, err
	}
	err = queryTx.Limit(p.Size).Offset(offset).Find(p.Data).Error
	p.Ok = err == nil
	return p, err
}

func crudAll(p *PaginationQ, queryTx *gorm.DB, list interface{}) (uint, error) {
	if p.Size < 1 {
		p.Size = 10
	}
	if p.Page < 1 {
		p.Page = 1
	}

	var total uint
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

// 随机生成指定位数的大写字母和数字的组合
func RandStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

//Make 加密方法
func makePassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

// //Check 检查方法
// func checkPassword(hashedPassword, password []byte) error {
// 	return bcrypt.CompareHashAndPassword(hashedPassword, password)
// }
