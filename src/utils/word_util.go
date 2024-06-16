package utils

import (
	"github.com/yanyiwu/gojieba"
)

func Cut(word string) []string {

	x := gojieba.NewJieba()
	defer x.Free()

	// x.AddWord("Markdown")
	// `AddWordEx` 支持指定词语的权重，作为 `AddWord` 权重太低加词失败的补充。
	// `tag` 参数可以为空字符串，也可以指定词性。
	// x.AddWordEx("比特币", 100000, "")

	return x.CutForSearch(word, true)

}
