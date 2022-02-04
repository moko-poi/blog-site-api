package config

import (
	"engineer-jobhunting-api/domain/model"
	"gorm.io/gorm"
)

// NewDB DBと接続する
func NewDB() *gorm.DB {
	db, err := gorm.Open("mysql", "user:password@tcp(sample_db)/sample?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(model.Blog{})

	return db
}
