package model

import "github.com/mouday/article-admin/src/utils"

type HistoryModel struct {
	Id          uint            `json:"historyId"`
	AssetId     uint            `json:"assetId"`
	UserId      string          `json:"userId"`
	Date        string          `json:"date"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Status      bool            `json:"status"`
	CreateTime  utils.LocalTime `gorm:"type:datetime;autoCreateTime" json:"createTime"`
	UpdateTime  utils.LocalTime `gorm:"type:datetime;autoUpdateTime" json:"updateTime"`
}

// 自定义表名
func (HistoryModel) TableName() string {
	return "tb_history"
}
