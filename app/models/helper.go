package models

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/pbkdf2"
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

//实现Django的加密
func encode(password string) string {
	algorithm := "pbkdf2_sha256" // 算法
	salt := []byte(RandStr(12))  // 盐，是一个随机字符串，每一个用户都不一样
	iterations := 260000         // 加密算法的迭代次数，15000 次
	digest := sha256.New         // digest 算法，使用 sha256
	fmt.Println("salt：" + string(salt))
	// 第一步：使用 pbkdf2 算法加密
	dk := pbkdf2.Key([]byte(password), salt, iterations, 32, digest)
	log.Println("dk：" + fmt.Sprintf("%x", dk))
	// log.Println("dk2：" + hex.EncodeToString(dk))

	// 第二步：Base64 编码
	base64 := base64.StdEncoding.EncodeToString(dk)
	log.Println("base64：" + base64)

	// 第三步：组合加密算法、迭代次数、盐、密码和分割符号 "$"
	pwd := fmt.Sprintf(
		"%s$%d$%s$%s",
		algorithm,
		iterations,
		string(salt),
		base64,
	)
	return string(pwd)
}

func decode(encoded string) map[string]string {
	str := "pbkdf2_sha256$216000$5uyVOJLx8oZK$cv6ayMeUUBiu8g8KXuBVyb+BVfdbiH8/FTGBoYBaICQ="
	split := strings.Split(str, "$")
	textMap := map[string]string{
		"algorithm":  split[0],
		"iterations": split[1],
		"salt":       split[2],
		"hash":       split[3],
	}
	return textMap
}
