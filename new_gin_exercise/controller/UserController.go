package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"mymod/common"
	"mymod/dto"
	"mymod/model"
	"mymod/response"
	"time"

	//"mymod/util"
	"mymod/utils"
)

type MyuserInfo struct {
	ID         uint
	Name       string
	Telephone  string
	Creattime  time.Time
	Updatetime time.Time
}

func Register(ctx *gin.Context) {
	DB := common.GetDb()
	//获取参数

	//var userInfo = model.User{}
	//json.NewDecoder(ctx.Request.Body).Decode(&userInfo)

	//name := userInfo.Name
	//telephone := userInfo.Telephone
	//password := userInfo.Password
	var requsetLoginUser = model.User{}

	ctx.Bind(&requsetLoginUser)
	name := requsetLoginUser.Name
	telephone := requsetLoginUser.Telephone
	password := requsetLoginUser.Password
	//数据验证

	if len(telephone) != 11 {
		response.Response(ctx, 422, 422, nil, "请填写正确的手机号码")
		return
	}

	if len(password) < 6 {
		response.Response(ctx, 422, 422, nil, "密码至少为六位数")
		return
	}

	if len(name) == 0 {
		name = utils.RandString(10)
	}

	//判断手机号是否正确
	if telephoneExit(DB, telephone) {

		response.Response(ctx, 422, 422, nil, "用户已经存在")

		return
	}

	//在这里进行密码的加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, 500, 500, nil, "加密错误")

		return
	}

	newUser := model.User{
		//Model:     gorm.Model{},
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)

	// 发送token
	token, err := common.ReleaseToken(newUser)

	if err != nil {

		response.Response(ctx, 422, 422, nil, "系统异常")
		log.Printf("token generate fail,%v", err)
		return
	}

	response.Success(ctx, gin.H{"token": token}, "注册成功")

	//如果名称没有传 就给一个随机的字符串
	//创建用户
}

func Requserlist(ctx *gin.Context) {
	DB := common.GetDb()
	// 这里要定义一个切片
	var userList []model.User
	var myUserList []MyuserInfo

	DB.Find(&userList)
	for _, value := range userList {
		fmt.Println(value)
		user := MyuserInfo{
			ID: value.Model.ID,
			//ID:         value.ID,
			Name:       value.Name,
			Telephone:  value.Telephone,
			Creattime:  value.CreatedAt,
			Updatetime: value.UpdatedAt,
		}
		myUserList = append(myUserList, user)
	}

	response.Success(ctx, gin.H{"user": myUserList}, "success")
}

func Login(ctx *gin.Context) {
	//获取参数
	DB := common.GetDb()
	// 在这里有好几种获取数据的方式

	// 使用map来获取请求体的参数
	//var requsetMap = make(map[string]string)
	//json.NewDecoder(ctx.Request.Body).Decode(&requsetMap)

	//通过结构体来获取请求头的参数
	var requsetUser = model.User{}
	//
	//json.NewDecoder(ctx.Request.Body).Decode(&requsetUser)
	//telephone := ctx.PostForm("telephone")
	//password := ctx.PostForm("password")
	//获取通过gin的框架自带的bind
	ctx.Bind(&requsetUser)
	fmt.Println(requsetUser)
	//
	telephone := requsetUser.Telephone
	password := requsetUser.Password

	fmt.Println(requsetUser)
	//password := ctx.PostForm("password")
	//telephone := ctx.PostForm("telephone")

	//数据验证
	if len(telephone) != 11 {
		response.Response(ctx, 422, 422, nil, "手机号码格式不正确")
		return
	}

	if len(password) < 6 {
		response.Response(ctx, 422, 422, nil, "密码至少为六位数")
		return
	}

	//判断手机号码是否存在
	var user model.User
	// 将这个结构体指针传入到函数中
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, 422, 422, nil, "用户不存在")
		return
	}

	//判断密码是否正确
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		response.Response(ctx, 422, 422, nil, "密码错误")
		return
	}

	// 发送token
	token, err := common.ReleaseToken(user)

	if err != nil {

		response.Response(ctx, 422, 422, nil, "系统异常")
		log.Printf("token generate fail,%v", err)
		return
	}

	response.Success(ctx, gin.H{"token": token, "user": dto.ToUserDto(user)}, "success")
}

func Test(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"code": 200,
		//"data":gin.H{"token":token},
		"msg": "success",
	})
}
func telephoneExit(db *gorm.DB, telephone string) bool {
	var user model.User
	// 将这个结构体指针传入到函数中
	db.Where("telephone = ?", telephone).First(&user)
	log.Println(user)
	if user.ID != 0 {
		return true
	}
	return false
}

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user") // 这里获取的user类型是接口 要用类型断言去判断
	fmt.Println(user.(model.User))
	ctx.JSON(200, gin.H{
		"user": dto.ToUserDto(user.(model.User)), // 这里应该是用到了类型断言
	})
}
