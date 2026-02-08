package dao

import (
	"log"
	"system/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMySQL() {
	var err error
	DB, err = gorm.Open(mysql.Open(
		"root:Paow7778200400@@tcp(127.0.0.1:3306)/winter_homework?charset=utf8mb4&parseTime=True&loc=Local",
	), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&models.User{}, &models.UserToken{}, &models.Homework{}, &models.Submission{})
	if err != nil {
		log.Fatal("自动迁移失败", err)
	}
}
