package common

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(110;not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

var DB *gorm.DB

func InitDb() *gorm.DB {

	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic("falied to connect database,err" + err.Error())
	}

	//db.AutoMigrate(&User{})
	db.AutoMigrate(&User{}) //自动迁移
	DB = db
	return db
}

func GetDb() *gorm.DB {
	return DB
}
