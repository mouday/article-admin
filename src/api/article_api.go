package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mouday/article-admin/src/config"
	"github.com/mouday/article-admin/src/dto"
	"github.com/mouday/article-admin/src/model"
	"github.com/mouday/article-admin/src/service"
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

	fmt.Println(row.Tags)

	if len(row.Tags) == 0 {
		row.Tags = service.ParseTags(row.Title)
	}

	service.AddTags(row.Tags)

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

	if len(row.Tags) == 0 {
		row.Tags = service.ParseTags(row.Title)
	}

	service.AddTags(row.Tags)

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
	params.Page = true

	pageVO := service.GetArticlePage(params)

	vo.Success(ctx, pageVO)
}

func RenderArticle(ctx *gin.Context) {

	params := &dto.PageDTO{}

	ctx.BindJSON(&params)

	pageVO := service.GetArticlePage(params)

	content := service.Render(pageVO)

	vo.Success(ctx, content)
}
