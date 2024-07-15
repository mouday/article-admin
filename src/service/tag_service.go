package service

import (
	"fmt"
	"strings"

	"github.com/mouday/article-admin/src/config"
	"github.com/mouday/article-admin/src/model"
)

func ParseTags(word string) []string {
	db := config.GetDB()
	// words := utils.Cut(word)

	tagList := []model.TagModel{}

	db.Model(&model.TagModel{}).Where("status = ?", true).Find(&tagList)

	var tags []string

	for _, row := range tagList {
		if strings.Contains(word, row.Title) {
			tags = append(tags, row.Title)
		}
	}

	fmt.Println(tags)

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
