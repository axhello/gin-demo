package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     "127.0.0.1",
		Port:     5432,
		User:     "mac",
		Password: "123456",
		DBName:   "coolpano",
	}
	return &dbConfig
}
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		// "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		"user=%s password=%s dbname=%s host=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Host,
		dbConfig.Port,
	)
}
