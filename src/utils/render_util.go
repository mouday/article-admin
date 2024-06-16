package utils

import (
	"github.com/flosch/pongo2/v6"
)

func Render(content string, data pongo2.Context) string {

	tpl, _ := pongo2.FromString(content)

	out, _ := tpl.Execute(data)

	return out
}
