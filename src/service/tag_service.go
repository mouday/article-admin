package service

import (
	"github.com/mouday/article-admin/src/config"
	"github.com/mouday/article-admin/src/model"
	"github.com/mouday/article-admin/src/utils"
)

func ParseTags(word string) []string {
	db := config.GetDB()
	words := utils.Cut(word)

	tagList := []model.TagModel{}
	db.Model(&model.TagModel{}).Where("status = ?", true).Where("title in ? ", words).Find(&tagList)

	var tags []string

	for _, row := range tagList {
		tags = append(tags, row.Title)
	}

	return tags
}

func AddTags(tags []string) {
	db := config.GetDB()

	for _, tag := range tags {
		row := &model.TagModel{}
		db.Model(&model.TagModel{}).FirstOrCreate(row, model.TagModel{
			Title: tag,
		})
	}
}
