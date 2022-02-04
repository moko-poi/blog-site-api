package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDB DBと接続する
func NewDB() *gorm.DB {
	dsn := "user:password@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
