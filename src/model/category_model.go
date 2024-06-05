package model

import "github.com/mouday/cron-admin/src/utils"

type CategoryModel struct {
	Id         uint            `json:"categoryId"`
	UserId     string          `json:"userId"`
	Title      string          `json:"title"`
	Status     bool            `json:"status"`
	CreateTime utils.LocalTime `gorm:"type:datetime;autoCreateTime" json:"createTime"`
	UpdateTime utils.LocalTime `gorm:"type:datetime;autoUpdateTime" json:"updateTime"`
}

// 自定义表名
func (CategoryModel) TableName() string {
	return "tb_category"
}
