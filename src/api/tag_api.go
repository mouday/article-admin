package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mouday/article-admin/src/config"
	"github.com/mouday/article-admin/src/dto"
	"github.com/mouday/article-admin/src/model"
	"github.com/mouday/article-admin/src/vo"
)

func AddTag(ctx *gin.Context) {
	row := &model.TagModel{}
	ctx.BindJSON(&row)

	db := config.GetDB()
	db.Model(&model.TagModel{}).Create(&row)

	vo.Success(ctx, row)
}

func UpdateTag(ctx *gin.Context) {
	row := &model.TagModel{}
	ctx.BindJSON(&row)

	db := config.GetDB()
	db.Model(&model.TagModel{}).Where("id = ?", row.Id).Updates(&row)

	vo.Success(ctx, nil)
}

func UpdateTagStatus(ctx *gin.Context) {
	params := &model.TagModel{}
	ctx.BindJSON(&params)

	db := config.GetDB()

	db.Model(&model.TagModel{}).Where("id = ?", params.Id).Update("status", params.Status)

	vo.Success(ctx, nil)

}

func DeleteTag(ctx *gin.Context) {
	row := &model.TagModel{}
	ctx.BindJSON(&row)

	db := config.GetDB()

	db.Where("id = ?", row.Id).Delete(&model.TagModel{})

	vo.Success(ctx, nil)
}

func GetTag(ctx *gin.Context) {

	row := &model.TagModel{}
	ctx.BindJSON(&row)

	db := config.GetDB()

	db.Model(&model.TagModel{}).Where("id = ?", row.Id).Find(&row)

	vo.Success(ctx, row)
}

func GetTagPage(ctx *gin.Context) {

	params := &dto.PageDTO{}

	ctx.BindJSON(&params)
	params.Page = true

	db := config.GetDB()

	TagList := []model.TagModel{}
	var count int64

	query := db.Model(&model.TagModel{})

	query.Count(&count)

	if count > 0 {
		query.Order("id desc").Find(&TagList)
	}

	vo.Success(ctx, vo.PageVO{
		List:  TagList,
		Total: count,
	})
}
