package model

import "github.com/mouday/cron-admin/src/utils"

type AssetModel struct {
	Id          uint            `json:"assetId"`
	UserId      string          `json:"userId"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Tags        string          `json:"tags"`
	CategoryId  int             `json:"categoryId"`
	Status      bool            `json:"status"`
	CreateTime  utils.LocalTime `gorm:"type:datetime;autoCreateTime" json:"createTime"`
	UpdateTime  utils.LocalTime `gorm:"type:datetime;autoUpdateTime" json:"updateTime"`
}

// 自定义表名
func (AssetModel) TableName() string {
	return "tb_asset"
}
