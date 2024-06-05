package model

import "github.com/mouday/cron-admin/src/utils"

type ContactModel struct {
	Id          uint            `json:"historyId"`
	UserId      string          `gorm:"index" json:"userId"`
	Date        string          `json:"date"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Status      bool            `json:"status"`
	CreateTime  utils.LocalTime `gorm:"type:datetime;autoCreateTime" json:"createTime"`
	UpdateTime  utils.LocalTime `gorm:"type:datetime;autoUpdateTime" json:"updateTime"`
}

// 自定义表名
func (ContactModel) TableName() string {
	return "tb_contact"
}
