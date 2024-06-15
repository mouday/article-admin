package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mouday/article-admin/src/config"
	"github.com/mouday/article-admin/src/dto"
	"github.com/mouday/article-admin/src/model"
	"github.com/mouday/article-admin/src/vo"
)

func AddCategory(ctx *gin.Context) {
	row := &model.CategoryModel{}
	ctx.BindJSON(&row)

	db := config.GetDB()
	db.Model(&model.CategoryModel{}).Create(&row)

	vo.Success(ctx, row)
}

func UpdateCategory(ctx *gin.Context) {
	row := &model.CategoryModel{}
	ctx.BindJSON(&row)

	db := config.GetDB()
	db.Model(&model.CategoryModel{}).Where("id = ?", row.Id).Updates(&row)

	vo.Success(ctx, nil)
}

func UpdateCategoryStatus(ctx *gin.Context) {
	params := &model.CategoryModel{}
	ctx.BindJSON(&params)

	db := config.GetDB()

	db.Model(&model.CategoryModel{}).Where("id = ?", params.Id).Update("status", params.Status)

	vo.Success(ctx, nil)

}

func DeleteCategory(ctx *gin.Context) {
	row := &model.CategoryModel{}
	ctx.BindJSON(&row)

	db := config.GetDB()

	db.Where("id = ?", row.Id).Delete(&model.CategoryModel{})

	vo.Success(ctx, nil)
}

func GetCategory(ctx *gin.Context) {

	row := &model.CategoryModel{}
	ctx.BindJSON(&row)

	db := config.GetDB()

	db.Model(&model.CategoryModel{}).Where("id = ?", row.Id).Find(&row)

	vo.Success(ctx, row)
}

func GetCategoryList(ctx *gin.Context) {

	params := &dto.PageDTO{}

	ctx.BindJSON(&params)

	db := config.GetDB()

	CategoryList := []model.CategoryModel{}
	var count int64

	query := db.Model(&model.CategoryModel{})

	query.Count(&count)

	if count > 0 {
		query.Order("id desc").Find(&CategoryList)
	}

	vo.Success(ctx, vo.PageVO{
		List:  CategoryList,
		Total: count,
	})
}
