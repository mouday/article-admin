package utils

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type DocData struct {
	Title string
	Url   string
}

// 获取文章数据
func GetArticleData(url string) DocData {
	// 示例HTML URL
	res, _ := http.Get(url)

	defer res.Body.Close()

	// 使用goquery解析HTML文档
	doc, _ := goquery.NewDocumentFromReader(res.Body)

	// 查询HTML文档
	title := doc.Find("#activity-name").First()

	data := DocData{
		Title: strings.TrimSpace(title.Text()),
		Url:   url,
	}

	return data
}
