package main

import (
	//"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)



type User struct {
	ID           int64
	Name *string `gorm:"default:'小王子'"`
	Age          int64
}
func main() {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err!= nil{
		panic(err)
	}
	defer db.Close()
	//db.AutoMigrate(&User{})
	//user := User{ID:1, Name: "q1mi", Age: 18}
	//user := User{Name: new(string), Age: 18}
	////var user = User{Age: 99}
	//db.Create(&user)

	// 根据主键查询第一条记录
	var user  User
	db.First(&user)

}
