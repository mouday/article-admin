package service

import (
	"github.com/mouday/article-admin/src/config"
	"github.com/mouday/article-admin/src/dto"
	"github.com/mouday/article-admin/src/model"
	"github.com/mouday/article-admin/src/vo"
)

func GetArticlePage(params *dto.PageDTO) vo.PageVO {
	db := config.GetDB()

	articleList := []model.ArticleModel{}
	var count int64

	query := db.Model(&model.ArticleModel{})

	if params.Keyword != "" {
		query = query.Where("title like ?", "%"+params.Keyword+"%").Or("tags like ?", "%\""+params.Keyword+"\"%")
	}

	if params.CategoryId >= 0 {
		query = query.Where("category_id = ?", params.CategoryId)
	}

	if params.StartDate != "" {
		query = query.Where("create_time >= ?", params.StartDate)
	}

	if params.EndDate != "" {
		query = query.Where("create_time <= ?", params.EndDate)
	}

	query.Count(&count)

	if count > 0 {

		if params.Page {
			query = query.Limit(params.GetSize()).Offset(params.PageOffset())
		}

		query.Order("id desc").Find(&articleList)

		// 查询文章的分类
		LoadCategory(articleList)
	}

	return vo.PageVO{
		List:  articleList,
		Total: count,
	}
}
