package models

import (
	"math/rand"
	"time"

	"github.com/jinzhu/gorm"
)

//PaginationQ gin handler query binding struct
type PaginationQ struct {
	Ok    bool        `json:"ok"`
	Size  uint        `form:"size" json:"size"`
	Page  uint        `form:"page" json:"page"`
	Data  interface{} `json:"data" comment:"muster be a pointer of slice gorm.Model"` // save pagination list
	Total uint        `json:"total"`
}

func GetTotal(p *PaginationQ, queryTx *gorm.DB, list interface{}) (uint, error) {
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
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < length; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
