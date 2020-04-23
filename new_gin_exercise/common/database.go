package common

import (
	"github.com/jinzhu/gorm"
	"mymod/model"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("falied to connect database,err" + err.Error())
	}

	//db.AutoMigrate(&User{})
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}

func GetDb() *gorm.DB {

	//DB = db
	return DB
}
