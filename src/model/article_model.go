package model

import "github.com/mouday/article-admin/src/utils"

type Tags []string

type ArticleModel struct {
	Id         uint            `json:"articleId"`
	Title      string          `json:"title"`
	Url        string          `json:"url"`
	Tags       Tags            `gorm:"type:varchar(255);serializer:json" json:"tags"`
	CategoryId uint            `json:"categoryId"`
	Category   CategoryModel   `gorm:"-" json:"category"`
	Status     bool            `json:"status"`
	CreateTime utils.LocalTime `gorm:"type:datetime;autoCreateTime" json:"createTime"`
	UpdateTime utils.LocalTime `gorm:"type:datetime;autoUpdateTime" json:"updateTime"`
}

// 自定义表名
func (ArticleModel) TableName() string {
	return "tb_article"
}

func (categoryAble *ArticleModel) GetCategoryId() uint {
	return categoryAble.CategoryId
}

func (categoryAble *ArticleModel) SetCategory(category CategoryModel) {
	categoryAble.Category = category
}
