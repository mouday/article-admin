package model

import "github.com/mouday/article-admin/src/utils"

type TagModel struct {
	Id         uint            `json:"tagId"`
	Title      string          `json:"title"`
	Status     bool            `json:"status"`
	CreateTime utils.LocalTime `gorm:"type:datetime;autoCreateTime" json:"createTime"`
	UpdateTime utils.LocalTime `gorm:"type:datetime;autoUpdateTime" json:"updateTime"`
}

// 自定义表名
func (TagModel) TableName() string {
	return "tb_tag"
}
