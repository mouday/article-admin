package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mouday/article-admin/src/config"
	"github.com/mouday/article-admin/src/dto"
	"github.com/mouday/article-admin/src/model"
	"github.com/mouday/article-admin/src/utils"
	"github.com/mouday/article-admin/src/vo"
)

func AddArticle(ctx *gin.Context) {
	row := &model.ArticleModel{}
	ctx.BindJSON(&row)

	if row.Title == "" {
		data := utils.GetArticleData(row.Url)
		row.Title = data.Title
	}

	db := config.GetDB()
	db.Model(&model.ArticleModel{}).Create(&row)

	vo.Success(ctx, row)
}

func UpdateArticle(ctx *gin.Context) {
	row := &model.ArticleModel{}
	ctx.BindJSON(&row)

	if row.Title == "" {
		data := utils.GetArticleData(row.Url)
		row.Title = data.Title
	}

	db := config.GetDB()
	db.Model(&model.ArticleModel{}).Where("id = ?", row.Id).Updates(&row)

	vo.Success(ctx, nil)
}

func UpdateArticleStatus(ctx *gin.Context) {
	params := &model.ArticleModel{}
	ctx.BindJSON(&params)

	db := config.GetDB()

	db.Model(&model.ArticleModel{}).Where("id = ?", params.Id).Update("status", params.Status)

	vo.Success(ctx, nil)

}

func DeleteArticle(ctx *gin.Context) {
	row := &model.ArticleModel{}
	ctx.BindJSON(&row)

	db := config.GetDB()

	db.Where("id = ?", row.Id).Delete(&model.ArticleModel{})

	vo.Success(ctx, nil)
}

func GetArticle(ctx *gin.Context) {

	row := &model.ArticleModel{}
	ctx.BindJSON(&row)

	db := config.GetDB()

	db.Model(&model.ArticleModel{}).Where("id = ?", row.Id).Find(&row)

	vo.Success(ctx, row)
}

func GetArticlePage(ctx *gin.Context) {

	params := &dto.PageDTO{}

	ctx.BindJSON(&params)

	db := config.GetDB()

	ArticleList := []model.ArticleModel{}
	var count int64

	query := db.Model(&model.ArticleModel{})

	query.Count(&count)

	if count > 0 {
		query.Order("id desc").Limit(params.GetSize()).Offset(params.PageOffset()).Find(&ArticleList)

		// 查询文章的分类
		categories := []model.CategoryModel{}
		var categoryIds []uint

		for _, row := range ArticleList {
			if row.CategoryId != 0 {
				categoryIds = append(categoryIds, row.CategoryId)
			}
		}

		if len(categoryIds) > 0 {
			db.Model(&model.CategoryModel{}).Where("id in ?", categoryIds).Find(&categories)

			categoryMap := make(map[uint]model.CategoryModel)
			for _, category := range categories {
				categoryMap[category.Id] = category
			}

			for index, _ := range ArticleList {
				ArticleList[index].Category = categoryMap[ArticleList[index].CategoryId]
			}
		}

	}

	vo.Success(ctx, vo.PageVO{
		List:  ArticleList,
		Total: count,
	})
}
