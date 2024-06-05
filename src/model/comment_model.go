package model

import "github.com/mouday/cron-admin/src/utils"

type CommentModel struct {
	Id         uint            `json:"commentId"`
	AssetId    uint            `json:"assetId"`
	UserId     string          `json:"userId"`
	Content    string          `json:"content"`
	Status     bool            `json:"status"`
	CreateTime utils.LocalTime `gorm:"type:datetime;autoCreateTime" json:"createTime"`
	UpdateTime utils.LocalTime `gorm:"type:datetime;autoUpdateTime" json:"updateTime"`
}

// 自定义表名
func (CommentModel) TableName() string {
	return "tb_comment"
}
