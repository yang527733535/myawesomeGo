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
	//"mymod/util"
	"mymod/utils"
)

func Register(ctx *gin.Context) {
	DB := common.GetDb()
	//获取参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	//数据验证

	if len(telephone) != 11 {
		ctx.JSON(422, gin.H{
			"msg": "请填写正确的手机号码",
		})
		return
	}

	if len(password) < 6 {
		ctx.JSON(422, gin.H{
			"msg": "密码至少为六位数",
		})
		return
	}

	if len(name) == 0 {
		name = utils.RandString(10)
	}

	//判断手机号是否正确
	if telephoneExit(DB, telephone) {
		ctx.JSON(422, gin.H{
			"msg": "用户已经存在",
		})
		return
	}

	//在这里进行密码的加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(500, gin.H{
			"code": 500,
			"msg":  "加密错误",
		})
		return
	}

	newUser := model.User{
		//Model:     gorm.Model{},
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)
	ctx.JSON(200, gin.H{
		"msg": "注册成功",
	})

	//如果名称没有传 就给一个随机的字符串
	//创建用户
}

func Login(ctx *gin.Context) {
	//获取参数
	DB := common.GetDb()
	password := ctx.PostForm("password")
	telephone := ctx.PostForm("telephone")

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

	//返回结果
	//ctx.JSON(200,gin.H{
	//	"code":200,
	//	"data":gin.H{"token":token},
	// 	"msg":"success",
	//})
	response.Success(ctx, gin.H{"token": token}, "success")
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
	user, _ := ctx.Get("user")
	fmt.Println(user.(model.User))
	ctx.JSON(200, gin.H{
		"user": dto.ToUserDto(user.(model.User)),
	})
}
