package main
import (
	"database/sql"
	"fmt"
	"time"

	//"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo 用户信息
type UserInfo struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	ID uint
	Name string
	Gender string
	Hobby string
}

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"` // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"` // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"` // 忽略本字段
}

func main() {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err!= nil{
		panic(err)
	}
	defer db.Close()
	////db.
	//db.AutoMigrate(&UserInfo{})
	db.AutoMigrate(&User{})
	//u1 := UserInfo{1, "七米", "S男", "篮球"}
	//u2 := UserInfo{2, "沙河娜扎", "女", "足球"}
	//// 创建记录
	//db.Create(&u1)
	//db.Create(&u2)


	// 查询
	var u = new(UserInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)
	var qimi UserInfo
	db.Find(&qimi, "hobby=?", "双色球")
	//
	//db.Model(&qimi).Update("hobby", "羽毛球")

	db.Delete(&qimi)
	//var uu UserInfo
	//db.Find(&uu, "hobby=?", "篮球")
	//fmt.Printf("%#v\n", uu)

	// 更新
	//db.Model(&uu).Update("hobby", "双色球")
}
