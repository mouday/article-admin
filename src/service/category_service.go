package service

import (
	"github.com/mouday/article-admin/src/config"
	"github.com/mouday/article-admin/src/model"
)

type CategoryAble interface {
	GetCategoryId() uint
	SetCategory(category model.CategoryModel)
}

func LoadCategory(list []model.ArticleModel) {
	db := config.GetDB()

	// 查询文章的分类
	categories := []model.CategoryModel{}
	var categoryIds []uint

	for _, row := range list {
		if row.GetCategoryId() != 0 {
			categoryIds = append(categoryIds, row.GetCategoryId())
		}
	}

	if len(categoryIds) > 0 {
		db.Model(&model.CategoryModel{}).Where("id in ?", categoryIds).Find(&categories)

		categoryMap := make(map[uint]model.CategoryModel)
		for _, category := range categories {
			categoryMap[category.Id] = category
		}

		for index, _ := range list {
			list[index].SetCategory(categoryMap[list[index].GetCategoryId()])
		}
	}
}
