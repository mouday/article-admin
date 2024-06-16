package service

import (
	"embed"

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

	// 读取文件内容
	content, _ := tmplFS.ReadFile("template/template.md")
	return utils.Render(string(content), pongo2.Context{
		"data": data,
	})
}
