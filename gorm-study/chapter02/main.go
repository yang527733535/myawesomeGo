package main

import (
	"fmt"
	//"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID   int64
	Name string `gorm:"default:'小王子'"`
	Age  int
}

func main() {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{})

	//u1 := User{Name: "DDFG", Age: 20}
	//
	////fmt.Println(db.NewRecord(user)) // 主键为空返回`true`
	//db.Create(&u1)

	var users []User
	var user User
	//db.Find(&users)
	//fmt.Println(users)
	//var users []struct{}
	//db.Where("name = ?", "xixixi").Find(&user)
	//db.Where("name <> ?", "xixixi").Find(&users)
	//db.Where("name IN (?)", []string{"q1mi", "xixixi"}).Find(&users)
	//db.Where("name LIKE ?", "%xi%").Find(&users)
	//db.Where("name = ? AND age = ?", "DDFG", "20").Find(&user)
	//db.Where(&User{Name: "DDFG", Age: 0}).First(&user)
	//db.Not([]int64{}).Find(&users)
	//db.Order("age desc, name").Find(&users)
	db.Order("age asc , name").Find(&users)
	//db.Where(map[string]interface{}{"name": "DDFG", "age": 20}).Find(&users)
	//db.Where([]int64{4, 8}).Find(&users)
	//fmt.Println(user)
	db.Where(User{Name: "sdf"}).FirstOrInit(&user)
	fmt.Println(users)
	fmt.Println(user)
	//fmt.Println(users)
	//user := User{Name: "q1mi", Age: 18}

	////fmt.Println(db.NewRecord(user)) // 主键为空返回`true`
	//db.Create(&user)
	//db.Create(&user2) // 创建user
	//fmt.Println(db.NewRecord(user)) // 创建`user`后返回`false`
}
