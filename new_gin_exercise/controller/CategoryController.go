package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"mymod/common"
	"mymod/model"
	"mymod/response"
	"strconv"
)

// 定义一个接口 用来控制这个模块的增删改查
type IcategoryController interface {
	RestController
}

//创建一个控制器实例 用这个结构体来实现上面的接口
type CategoryController struct {
	DB *gorm.DB
}

func NewCategoryController() IcategoryController {
	db := common.GetDb()
	//db.AutoMigrate()
	db.AutoMigrate(model.Category{}) // 这里有一个问题
	//CategoryController := CategoryController{DB: db}
	return CategoryController{DB: db}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var requestCategory model.Category
	ctx.Bind(&requestCategory)

	fmt.Println(requestCategory)
	//ctx.Bind(&requestCategory)
	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "name是必须的")
		return
	}
	c.DB.Create(&requestCategory)
	response.Success(ctx, gin.H{"catagort": requestCategory}, "创建成功")
}

func (c CategoryController) Update(ctx *gin.Context) {
	var requestCategory model.Category
	ctx.Bind(&requestCategory)
	fmt.Println(requestCategory)
	if requestCategory.Name == "" {
		response.Fail(ctx, nil, "name是必须的")
		return
	}
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	fmt.Println(categoryId)

	var updateCategory model.Category
	if c.DB.First(&updateCategory, categoryId).RecordNotFound() {
		response.Fail(ctx, nil, "分类不存在")
		return
	}
	c.DB.Model(&updateCategory).Update("name", requestCategory.Name)
	response.Success(ctx, gin.H{"category": updateCategory}, "修改成功")
}

func (c CategoryController) Show(ctx *gin.Context) {

	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	var Category model.Category
	if c.DB.First(&Category, categoryId).RecordNotFound() {
		response.Fail(ctx, nil, "分类不存在")
		return
	}
	response.Success(ctx, gin.H{"category": Category}, "成功")
}

func (c CategoryController) Delete(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Params.ByName("id"))
	//var Category model.Category

	if err := c.DB.Delete(model.Category{}, categoryId).Error; err != nil {
		response.Fail(ctx, nil, "删除失败")
		return
	}

	response.Success(ctx, nil, "删除成功")
}
