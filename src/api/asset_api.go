package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mouday/cron-admin/src/config"
	"github.com/mouday/cron-admin/src/form"
	"github.com/mouday/cron-admin/src/model"
	"github.com/mouday/cron-admin/src/vo"
)

type AssetForm struct {
	Id          uint   `json:"assetId"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
	CategoryId  int    `json:"categoryId"`
	Status      bool   `json:"status"`
}

func AddAsset(ctx *gin.Context) {
	form := &AssetForm{}
	ctx.BindJSON(&form)

	row := &model.AssetModel{
		Title:       form.Title,
		Description: form.Description,
		Tags:        form.Tags,
		CategoryId:  form.CategoryId,
		Status:      form.Status,
	}

	db := config.GetDB()
	db.Model(&model.AssetModel{}).Create(&row)

	vo.Success(ctx, row)

}

func UpdateAsset(ctx *gin.Context) {
	row := &model.AssetModel{}
	ctx.BindJSON(&row)

	db := config.GetDB()
	db.Model(&model.AssetModel{}).Where("id = ?", row.Id).Updates(&row)

	vo.Success(ctx, nil)
}

func UpdateAssetStatus(ctx *gin.Context) {
	params := &model.AssetModel{}
	ctx.BindJSON(&params)

	db := config.GetDB()

	db.Model(&model.AssetModel{}).Where("id = ?", params.Id).Update("status", params.Status)

	vo.Success(ctx, nil)

}

func RemoveAsset(ctx *gin.Context) {
	row := &model.AssetModel{}
	ctx.BindJSON(&row)

	db := config.GetDB()

	db.Where("id = ?", row.Id).Delete(&model.AssetModel{})

	vo.Success(ctx, nil)
}

func GetAsset(ctx *gin.Context) {

	row := &model.AssetModel{}
	ctx.BindJSON(&row)

	db := config.GetDB()

	db.Model(&model.AssetModel{}).Where("id = ?", row.Id).Find(&row)

	vo.Success(ctx, row)
}

func GetAssetList(ctx *gin.Context) {

	params := &form.PageForm{}

	ctx.BindJSON(&params)

	db := config.GetDB()

	AssetList := []model.AssetModel{}
	var count int64

	query := db.Model(&model.AssetModel{})

	if params.UserId != 0 {
		query.Where("user_id = ?", params.UserId)
	}

	query.Count(&count)

	query.Order("id desc").Limit(params.GetSize()).Offset(params.PageOffset()).Find(&AssetList)

	vo.Success(ctx, gin.H{
		"list":  AssetList,
		"total": count,
	})
}
