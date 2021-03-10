package helper

import (
	"errors"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// SecretKey 私钥
var SecretKey = os.Getenv("SECRET_KEY")

// UserClaims struct
type UserClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenerateToken 生成Token
func GenerateToken(username string, expireDuration time.Duration) (string, error) {
	expire := time.Now().Add(expireDuration)
	// 将 uid，用户角色， 过期时间作为数据写入 token 中
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expire.Unix(),
		},
	})
	// SecretKey 用于对用户数据进行签名，不能暴露
	return token.SignedString([]byte(SecretKey))
}

// ParseToken 解析JWT
func ParseToken(tokenStr string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
