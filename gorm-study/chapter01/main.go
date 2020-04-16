package main

import (
	"fmt"
	"time"

	//"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo 用户信息

type Userinfo struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	ID        uint
	Name      string
	Gender    string
	Hobby     string
}

func main() {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	fmt.Println("启动成功")
	defer db.Close()

	db.AutoMigrate(&Userinfo{})

	// 自动迁移
	//u1 := Userinfo{
	//	ID:     1,
	//	Gender: "man",
	//	Name:   "xix",
	//	Hobby:  "ball",
	//}
	////fmt.Println(u1)
	//db.Create(&u1)
	//查询
	//var u = new(Userinfo)
	//db.First(u)
	//fmt.Println(u)
	//查找第一条记录

	//fmt.Println(u)
	var uu Userinfo
	////////条件查询 每次查询的结果 要先创建一个变量来接受
	db.Find(&uu, "hobby=?", "ball")
	fmt.Println(uu)
	//
	//// 更新
	//db.Model(&uu).Update("hobby", "ball2")

	////删除
	db.Delete(&uu)
	//var u = new(UserInfo)
	////u := &UserInfo{}
	//db.First(u)
	//fmt.Println(u)
}
