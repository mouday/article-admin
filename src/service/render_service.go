package service

import (
	"embed"
	"sort"

	"github.com/flosch/pongo2/v6"
	"github.com/mouday/article-admin/src/model"
	"github.com/mouday/article-admin/src/utils"
	"github.com/mouday/article-admin/src/vo"
)

//go:embed template/*
var tmplFS embed.FS

type RenderData struct {
	Category model.CategoryModel
	Children []model.ArticleModel
}

type RenderDataList []RenderData

// 实现sort.Interface的Len方法
func (s RenderDataList) Len() int {
	return len(s)
}

// 实现sort.Interface的Swap方法
func (s RenderDataList) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// 实现sort.Interface的Less方法，根据年龄排序
func (s RenderDataList) Less(i, j int) bool {
	if s[i].Category.OrderValue == s[j].Category.OrderValue {
		return s[i].Category.Id > s[j].Category.Id
	} else {
		return s[i].Category.OrderValue > s[j].Category.OrderValue
	}
}

func Render(pageVO vo.PageVO) string {
	data := make(map[uint]RenderData)

	for _, row := range pageVO.List.([]model.ArticleModel) {

		_, ok := data[row.CategoryId]
		if ok {
			children := append(data[row.CategoryId].Children, row)

			data[row.CategoryId] = RenderData{
				Category: row.Category,
				Children: children,
			}

		} else {

			data[row.CategoryId] = RenderData{
				Category: row.Category,
				Children: []model.ArticleModel{row},
			}
		}
	}

	renderDataList := RenderDataList{}

	for _, value := range data {
		renderDataList = append(renderDataList, value)
	}

	sort.Stable(renderDataList)

	// 读取文件内容
	content, _ := tmplFS.ReadFile("template/template.md")
	return utils.Render(string(content), pongo2.Context{
		"rows": renderDataList,
	})
}
