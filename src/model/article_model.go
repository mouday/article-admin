package model

import "github.com/mouday/article-admin/src/utils"

type ArticleModel struct {
	Id         uint            `json:"articleId"`
	Title      string          `json:"title"`
	Url        string          `json:"url"`
	Tags       string          `json:"tags"`
	CategoryId int             `json:"categoryId"`
	Status     bool            `json:"status"`
	CreateTime utils.LocalTime `gorm:"type:datetime;autoCreateTime" json:"createTime"`
	UpdateTime utils.LocalTime `gorm:"type:datetime;autoUpdateTime" json:"updateTime"`
}

// 自定义表名
func (ArticleModel) TableName() string {
	return "tb_article"
}
