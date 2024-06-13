package utils

import (
	"fmt"
	"testing"
)

func TestGetArticleData(t *testing.T) {
	url := "https://mp.weixin.qq.com/s/gnMItmWB1ZFZwzZex7f0RQ"
	data := GetArticleData(url)

	// 打印结果
	fmt.Println("title:", data.Title)
	fmt.Println("url:", data.Url)

}
